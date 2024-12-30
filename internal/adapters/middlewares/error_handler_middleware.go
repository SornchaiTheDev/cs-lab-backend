package middlewares

import (
	"errors"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var e *cserrors.Error
	if errors.As(err, &e) {
		return c.Status(int(e.Code)).JSON(fiber.Map{
			"code":  e.Code,
			"error": e.Message,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":  "INTERNAL_SERVER_ERROR",
		"error": "Internal Server Error",
	})
}
