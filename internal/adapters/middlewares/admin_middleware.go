package middlewares

import (
	"github.com/SornchaiTheDev/cs-lab-backend/constants"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	for _, role := range user.Roles {
		if role == constants.ADMIN {
			return c.Next()
		}
	}

	return cserrors.New(cserrors.UNAUTHORIZED, "Unauthorized")
}
