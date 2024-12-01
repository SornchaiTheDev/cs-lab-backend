package rerror

import "github.com/gofiber/fiber/v2"

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}
