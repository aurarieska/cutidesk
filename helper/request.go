package helper

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"cutidesk/app/model"
)

const (
	DefaultLimit = 10
	MaxLimit     = 100
	maxSearchLen = 100
)

func RequestContext(c *fiber.Ctx) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.UserContext(), 5*time.Second)
}

func ParamID(c *fiber.Ctx) (int, error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return 0, BadRequest("id harus berupa angka positif")
	}
	return id, nil
}

func ParsePage(c *fiber.Ctx) (model.Page, error) {
	limit := c.QueryInt("limit", DefaultLimit)
	if limit < 1 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}

	page := model.Page{Limit: limit}
	if raw := c.Query("cursor"); raw != "" {
		cursor, err := DecodeCursor(raw)
		if err != nil {
			return model.Page{}, BadRequest("cursor tidak valid")
		}
		page.After = &cursor
	}
	return page, nil
}

func QuerySearch(c *fiber.Ctx) (string, error) {
	search := strings.TrimSpace(c.Query("search"))
	if len(search) > maxSearchLen {
		return "", BadRequest("search maksimal 100 karakter")
	}
	return search, nil
}

func QueryBool(c *fiber.Ctx, key string) (*bool, error) {
	raw := c.Query(key)
	if raw == "" {
		return nil, nil
	}
	v, err := strconv.ParseBool(raw)
	if err != nil {
		return nil, BadRequest(key + " harus true atau false")
	}
	return &v, nil
}

func QueryPositiveInt(c *fiber.Ctx, key string) (*int, error) {
	raw := c.Query(key)
	if raw == "" {
		return nil, nil
	}
	v, err := strconv.Atoi(raw)
	if err != nil || v < 1 {
		return nil, BadRequest(key + " harus berupa angka positif")
	}
	return &v, nil
}

func QueryDate(c *fiber.Ctx, key string) (string, error) {
	raw := c.Query(key)
	if raw == "" {
		return "", nil
	}
	if _, err := time.Parse("2006-01-02", raw); err != nil {
		return "", BadRequest(key + " harus berformat YYYY-MM-DD")
	}
	return raw, nil
}

func QueryEnum(c *fiber.Ctx, key string, allowed ...string) (string, error) {
	raw := strings.TrimSpace(c.Query(key))
	if raw == "" {
		return "", nil
	}
	for _, a := range allowed {
		if raw == a {
			return raw, nil
		}
	}
	return "", BadRequest(key + " harus salah satu dari: " + strings.Join(allowed, ", "))
}