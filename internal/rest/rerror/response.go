package rerror

import (
	"github.com/gofiber/fiber/v2"
)

func Res(c *fiber.Ctx, err RespError, msg string) error {
	code := MapErrorWithFiberStatus(err)

	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}
