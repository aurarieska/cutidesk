INSERT INTO roles (name, description) VALUES
    ('admin',    'Mengelola data pengguna dan master jenis cuti'),
    ('hr',       'Melihat seluruh pengajuan cuti serta menyetujui atau menolaknya'),
    ('employee', 'Mengajukan cuti dan melihat riwayat pengajuan miliknya sendiri')
ON CONFLICT (name) DO NOTHING;

INSERT INTO permissions (name, description) VALUES
    ('user:list',             'Melihat daftar pengguna'),
    ('user:read',             'Melihat detail pengguna'),
    ('user:create',           'Membuat pengguna baru'),
    ('user:update',           'Mengubah profil pengguna'),
    ('user:deactivate',       'Mengaktifkan atau menonaktifkan akun pengguna'),
    ('role:assign',           'Mengubah role milik user lain'),
    ('leave-type:read',       'Melihat jenis cuti'),
    ('leave-type:create',     'Menambah jenis cuti'),
    ('leave-type:update',     'Mengubah jenis cuti'),
    ('leave-type:deactivate', 'Mengaktifkan atau menonaktifkan jenis cuti'),
    ('leave-request:create',  'Mengajukan cuti'),
    ('leave-request:read:any','Melihat pengajuan cuti milik siapa pun'),
    ('leave-request:decide',  'Menyetujui atau menolak pengajuan cuti')
ON CONFLICT (name) DO NOTHING;

INSERT INTO role_permissions (role_name, permission_name) VALUES
    ('admin', 'user:list'),
    ('admin', 'user:read'),
    ('admin', 'user:create'),
    ('admin', 'user:update'),
    ('admin', 'user:deactivate'),
    ('admin', 'role:assign'),
    ('admin', 'leave-type:read'),
    ('admin', 'leave-type:create'),
    ('admin', 'leave-type:update'),
    ('admin', 'leave-type:deactivate'),
    ('hr', 'leave-type:read'),
    ('hr', 'leave-request:read:any'),
    ('hr', 'leave-request:decide'),
    ('employee', 'leave-type:read'),
    ('employee', 'leave-request:create')
ON CONFLICT DO NOTHING;

INSERT INTO leave_types (name, description) VALUES
    ('Cuti Tahunan',  'Cuti tahunan reguler'),
    ('Cuti Sakit',    'Cuti karena sakit'),
    ('Cuti Penting',  'Cuti untuk keperluan keluarga atau urusan penting')
ON CONFLICT DO NOTHING;