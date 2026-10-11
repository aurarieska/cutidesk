package helper

import (
	"encoding/csv"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"cutidesk/app/model"
)

const (
	FormatJSON = fiber.MIMEApplicationJSON
	FormatCSV  = "text/csv"
)

func Negotiate(c *fiber.Ctx, offered ...string) (string, error) {
	accept := strings.TrimSpace(c.Get(fiber.HeaderAccept))
	if accept == "" {
		return offered[0], nil
	}
	chosen := c.Accepts(offered...)
	if chosen == "" {
		return "", NotAcceptable(
			"format yang diminta tidak tersedia, pilih salah satu dari: " + strings.Join(offered, ", "))
	}
	return chosen, nil
}

func csvSafe(s string) string {
	if s == "" {
		return s
	}
	switch s[0] {
	case '=', '+', '-', '@', '\t', '\r':
		return "'" + s
	}
	return s
}

func writeCSV(c *fiber.Ctx, filename string, header []string, rows [][]string) error {
	c.Set(fiber.HeaderContentType, FormatCSV+"; charset=utf-8")
	c.Set(fiber.HeaderContentDisposition, `attachment; filename="`+filename+`"`)

	var buffer strings.Builder
	w := csv.NewWriter(&buffer)
	if err := w.Write(header); err != nil {
		return Internal(err)
	}
	for _, row := range rows {
		safe := make([]string, len(row))
		for i, cell := range row {
			safe[i] = csvSafe(cell)
		}
		if err := w.Write(safe); err != nil {
			return Internal(err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return Internal(err)
	}
	return c.SendString(buffer.String())
}

func WriteUsersCSV(c *fiber.Ctx, users []model.User) error {
	rows := make([][]string, 0, len(users))
	for _, u := range users {
		rows = append(rows, []string{
			strconv.Itoa(u.ID), u.Username, u.FullName, u.Email, u.Role,
			strconv.FormatBool(u.IsActive), u.CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
		})
	}
	return writeCSV(c, "users.csv",
		[]string{"id", "username", "full_name", "email", "role", "is_active", "created_at"}, rows)
}

func WriteLeaveRequestsCSV(c *fiber.Ctx, items []model.LeaveRequest) error {
	rows := make([][]string, 0, len(items))
	for _, lr := range items {
		decidedBy, decidedAt, note := "", "", ""
		if lr.Decision != nil {
			decidedBy = lr.Decision.DecidedBy.Username
			decidedAt = lr.Decision.DecidedAt.UTC().Format("2006-01-02T15:04:05Z")
			note = lr.Decision.Note
		}
		rows = append(rows, []string{
			strconv.Itoa(lr.ID), lr.Employee.Username, lr.Employee.FullName, lr.LeaveType.Name,
			lr.StartDate, lr.EndDate, lr.Reason, lr.Status,
			decidedBy, decidedAt, note, lr.CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
		})
	}
	return writeCSV(c, "leave-requests.csv", []string{
		"id", "employee", "employee_name", "leave_type", "start_date", "end_date",
		"reason", "status", "decided_by", "decided_at", "decision_note", "created_at",
	}, rows)
}