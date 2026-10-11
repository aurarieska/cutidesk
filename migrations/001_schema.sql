-- =====================================================================
-- 001_schema.sql — skema final CutiDesk
-- Dijalankan SEKALI pada database kosong. Aman diulang (IF NOT EXISTS).
-- =====================================================================

CREATE EXTENSION IF NOT EXISTS btree_gist;

-- RBAC
CREATE TABLE IF NOT EXISTS roles (
    name        VARCHAR(20)  PRIMARY KEY,
    description VARCHAR(150) NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS permissions (
    name        VARCHAR(50)  PRIMARY KEY,
    description VARCHAR(150) NOT NULL
);

CREATE TABLE IF NOT EXISTS role_permissions (
    role_name       VARCHAR(20) NOT NULL REFERENCES roles(name)       ON DELETE CASCADE,
    permission_name VARCHAR(50) NOT NULL REFERENCES permissions(name) ON DELETE CASCADE,
    PRIMARY KEY (role_name, permission_name)
);

-- users
CREATE TABLE IF NOT EXISTS users (
    id         SERIAL       PRIMARY KEY,
    username   VARCHAR(30)  NOT NULL,
    email      VARCHAR(120) NOT NULL,
    full_name  VARCHAR(100) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    role       VARCHAR(20)  NOT NULL REFERENCES roles(name) ON UPDATE CASCADE,
    is_active  BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS users_username_lower_key ON users (LOWER(username));
CREATE UNIQUE INDEX IF NOT EXISTS users_email_lower_key    ON users (LOWER(email));
CREATE INDEX        IF NOT EXISTS users_role_idx           ON users (role);
CREATE INDEX        IF NOT EXISTS users_created_at_id_idx  ON users (created_at DESC, id DESC);

-- refresh_tokens
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id         BIGSERIAL   PRIMARY KEY,
    user_id    INTEGER     NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash TEXT        NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS refresh_tokens_user_id_idx ON refresh_tokens (user_id);

-- leave_types
CREATE TABLE IF NOT EXISTS leave_types (
    id          SERIAL       PRIMARY KEY,
    name        VARCHAR(50)  NOT NULL,
    description VARCHAR(200) NOT NULL DEFAULT '',
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX IF NOT EXISTS leave_types_name_lower_key    ON leave_types (LOWER(name));
CREATE INDEX        IF NOT EXISTS leave_types_created_at_id_idx ON leave_types (created_at DESC, id DESC);

-- leave_requests
CREATE TABLE IF NOT EXISTS leave_requests (
    id            SERIAL       PRIMARY KEY,
    employee_id   INTEGER      NOT NULL REFERENCES users(id),
    leave_type_id INTEGER      NOT NULL REFERENCES leave_types(id),
    start_date    DATE         NOT NULL,
    end_date      DATE         NOT NULL,
    reason        VARCHAR(500) NOT NULL,
    status        VARCHAR(10)  NOT NULL DEFAULT 'pending',
    decided_by    INTEGER      REFERENCES users(id),
    decided_at    TIMESTAMPTZ,
    decision_note VARCHAR(500),
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    CONSTRAINT leave_requests_status_chk CHECK (status IN ('pending', 'approved', 'rejected')),
    CONSTRAINT leave_requests_dates_chk  CHECK (end_date >= start_date),
    CONSTRAINT leave_requests_decision_chk CHECK (
        (status = 'pending' AND decided_by IS NULL AND decided_at IS NULL)
        OR
        (status <> 'pending' AND decided_by IS NOT NULL AND decided_at IS NOT NULL)
    ),
    CONSTRAINT leave_requests_no_overlap EXCLUDE USING gist (
        employee_id WITH =,
        daterange(start_date, end_date, '[]') WITH &&
    ) WHERE (status IN ('pending', 'approved'))
);

CREATE INDEX IF NOT EXISTS leave_requests_created_at_id_idx
    ON leave_requests (created_at DESC, id DESC);
CREATE INDEX IF NOT EXISTS leave_requests_employee_created_idx
    ON leave_requests (employee_id, created_at DESC, id DESC);
CREATE INDEX IF NOT EXISTS leave_requests_status_created_idx
    ON leave_requests (status, created_at DESC, id DESC);
CREATE INDEX IF NOT EXISTS leave_requests_leave_type_idx
    ON leave_requests (leave_type_id);

-- leave_request_status_history
CREATE TABLE IF NOT EXISTS leave_request_status_history (
    id               BIGSERIAL    PRIMARY KEY,
    leave_request_id INTEGER      NOT NULL REFERENCES leave_requests(id) ON DELETE CASCADE,
    from_status      VARCHAR(10),
    to_status        VARCHAR(10)  NOT NULL,
    changed_by       INTEGER      NOT NULL REFERENCES users(id),
    note             VARCHAR(500),
    created_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT leave_history_to_chk   CHECK (to_status IN ('pending', 'approved', 'rejected')),
    CONSTRAINT leave_history_from_chk CHECK (from_status IS NULL OR from_status IN ('pending', 'approved', 'rejected'))
);
CREATE INDEX IF NOT EXISTS leave_history_request_idx
    ON leave_request_status_history (leave_request_id, created_at, id);