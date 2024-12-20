package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	// user := c.Locals("user").(*auth.JWTClaims)
	// fmt.Println(user.Roles)
	return c.Next()
}
