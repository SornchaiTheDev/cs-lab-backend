package rest

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/infrastructure/auth"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/rerror"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthRouter(router fiber.Router, appConfig *configs.Config, userService services.UserService) {
	authRouter := router.Group("/auth")

	googleAuth := auth.NewGoogleAuth(appConfig)

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

		user, err := userService.GetByEmail(userInfo.Email)
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

		user, err := userService.GetByUsername(credential.Username)
		if err != nil {
			return rerror.Unauthorized(c)
		}

		password, err := userService.GetPasswordByID(user.ID)
		if err != nil {
			return rerror.Unauthorized(c)
		}

		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(credential.Password))
		if err != nil {
			return rerror.Unauthorized(c)
		}

		token, err := auth.SignJWT(user, appConfig.JWTSecret)
		if err != nil {
			return rerror.InternalServerError(c)
		}

		return c.JSON(fiber.Map{
			"message": "OK",
			"token":   token,
		})
	})

}
