package model

import "time"

type LeaveType struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateLeaveTypeRequest struct {
	Name        string `json:"name"        validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=200"`
}

type ReplaceLeaveTypeRequest struct {
	Name        string `json:"name"        validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=200"`
}

type PatchLeaveTypeRequest struct {
	Name        *string `json:"name,omitempty"        validate:"omitnil,min=3,max=50"`
	Description *string `json:"description,omitempty" validate:"omitnil,max=200"`
}

type LeaveTypeFilter struct {
	Page
	Search   string
	IsActive *bool
}