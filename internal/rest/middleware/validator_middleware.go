package middleware

import (
	"github.com/SornchaiTheDev/cs-lab-backend/internal/validator"
	"github.com/gofiber/fiber/v2"
)

func ValidateMiddleware(r any) func(*fiber.Ctx) error {
	appValidator := validator.NewAppValidator()

	return func(c *fiber.Ctx) error {
		c.BodyParser(&r)

		if errs := appValidator.Validate(r); len(errs) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Bad Request",
				"fields": errs,
			})
		}

		c.Locals("request", r)

		return c.Next()
	}
}
