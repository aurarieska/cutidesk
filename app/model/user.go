package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"  validate:"required,min=3,max=30,username"`
	Email    string `json:"email"     validate:"required,email,max=120"`
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
	Password string `json:"password"  validate:"required,max=72,strongpassword"`
	Role     string `json:"role"      validate:"required"`
}

type ReplaceUserRequest struct {
	Email    string `json:"email"     validate:"required,email,max=120"`
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
}

type PatchUserRequest struct {
	Email    *string `json:"email,omitempty"     validate:"omitnil,email,max=120"`
	FullName *string `json:"full_name,omitempty" validate:"omitnil,min=3,max=100"`
}

type StatusRequest struct {
	IsActive *bool `json:"is_active" validate:"required"`
}

type AssignRoleRequest struct {
	Role string `json:"role" validate:"required"`
}

type UserFilter struct {
	Page
	Search   string
	Role     string
	IsActive *bool
}