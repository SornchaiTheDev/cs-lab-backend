package rest

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/infrastructure/auth"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/gofiber/fiber/v2"
)

func NewAuthRouter(router fiber.Router, c *configs.Config, userService services.UserService) {
	authRouter := router.Group("/auth")

	googleAuth := auth.NewGoogleAuth(c)

	// Google OAuth2
	authRouter.Get("/sign-in/google", func(c *fiber.Ctx) error {
		url, err := googleAuth.GenerateAuthURL()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Redirect(url)
	},
	)

	authRouter.Get("/sign-in/google/callback", func(c *fiber.Ctx) error {
		state := c.Query("state")
		if !googleAuth.VerifyState(state) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"code": "INVALID_STATE"})
		}

		ctx := context.Background()

		code := c.Query("code")

		userInfo, err := googleAuth.GetUserInfo(ctx, code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code": "EXCHANGE_FAILED",
			})
		}

		user, err := userService.GetUserByEmail(userInfo.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code": "USER_NOT_FOUND",
				"err":  err.Error(),
			})

		}

		return c.JSON(user)
	})

	authRouter.Post("/sign-in/credential", ValidateMiddleware(&requests.Credential{}), func(c *fiber.Ctx) error {
		credential := c.Locals("request").(*requests.Credential)

		return c.JSON(fiber.Map{
			"message":  "OK",
			"password": credential.Password,
		})
	})

}
