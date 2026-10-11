package model

import "time"

type PersonRef struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type LeaveTypeRef struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Decision struct {
	DecidedBy PersonRef `json:"decided_by"`
	DecidedAt time.Time `json:"decided_at"`
	Note      string    `json:"note,omitempty"`
}

type LeaveRequest struct {
	ID        int          `json:"id"`
	Employee  PersonRef    `json:"employee"`
	LeaveType LeaveTypeRef `json:"leave_type"`
	StartDate string       `json:"start_date"`
	EndDate   string       `json:"end_date"`
	Reason    string       `json:"reason"`
	Status    string       `json:"status"`
	Decision  *Decision    `json:"decision"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"
)

type SubmitLeaveRequest struct {
	LeaveTypeID int    `json:"leave_type_id" validate:"required,min=1"`
	StartDate   string `json:"start_date"    validate:"required,datetime=2006-01-02"`
	EndDate     string `json:"end_date"      validate:"required,datetime=2006-01-02"`
	Reason      string `json:"reason"        validate:"required,min=5,max=500"`
}

type NewLeaveRequest struct {
	EmployeeID  int
	LeaveTypeID int
	StartDate   string
	EndDate     string
	Reason      string
}

type DecideLeaveRequest struct {
	Status string `json:"status" validate:"required,oneof=approved rejected"`
	Note   string `json:"note"   validate:"max=500"`
}

type StatusHistory struct {
	ID             int64     `json:"id"`
	LeaveRequestID int       `json:"leave_request_id"`
	FromStatus     *string   `json:"from_status"`
	ToStatus       string    `json:"to_status"`
	ChangedBy      PersonRef `json:"changed_by"`
	Note           string    `json:"note,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

type LeaveRequestFilter struct {
	Page
	EmployeeID  *int
	LeaveTypeID *int
	Status      string
	StartFrom   string
	StartTo     string
}