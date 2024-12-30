package middlewares

import (
	"github.com/SornchaiTheDev/cs-lab-backend/constants"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/SornchaiTheDev/cs-lab-backend/infrastructure/auth"
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.JWTClaims)

	for _, role := range user.Roles {
		if role == constants.ADMIN {
			return c.Next()
		}
	}

	return &cserrors.Error{Code: cserrors.UNAUTHORIZED, Message: "Unauthorized"}
}
