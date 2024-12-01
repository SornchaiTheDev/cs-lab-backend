package rerror

import "github.com/gofiber/fiber/v2"

func BadRequest(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Bad Request",
	})
}
