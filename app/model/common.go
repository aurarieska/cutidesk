package model

import "time"

type WebResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    any         `json:"data,omitempty"`
	Meta    *CursorMeta `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Success   bool              `json:"success"`
	Code      string            `json:"code"`
	Message   string            `json:"message"`
	Fields    map[string]string `json:"fields,omitempty"`
	RequestID string            `json:"request_id,omitempty"`
}

type CursorMeta struct {
	Limit      int    `json:"limit"`
	NextCursor string `json:"next_cursor,omitempty"`
	HasMore    bool   `json:"has_more"`
}

type Cursor struct {
	CreatedAt time.Time
	ID        int
}

type Page struct {
	Limit int
	After *Cursor
}