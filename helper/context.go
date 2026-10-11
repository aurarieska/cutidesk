package helper

import (
	"github.com/gofiber/fiber/v2"

	"cutidesk/app/model"
)

const LocalsAuthUser = "authUser"

func CurrentUser(c *fiber.Ctx) (model.AuthUser, bool) {
	user, ok := c.Locals(LocalsAuthUser).(model.AuthUser)
	return user, ok
}

func RequireUser(c *fiber.Ctx) (model.AuthUser, error) {
	user, ok := CurrentUser(c)
	if !ok {
		return model.AuthUser{}, Unauthorized("belum terautentikasi")
	}
	return user, nil
}

func RequestID(c *fiber.Ctx) string {
	id, _ := c.Locals("requestid").(string)
	return id
}