package rest

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/infrastructure/auth"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/middlewares"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// TODO: redirect back to frontend and set cookie
func NewAuthRouter(router fiber.Router, appConfig *configs.Config, userService services.UserService) {
	authRouter := router.Group("/auth")

	googleAuth := auth.NewGoogleAuth(appConfig)

	// Google OAuth2
	authRouter.Get("/sign-in/google", func(c *fiber.Ctx) error {
		url, err := googleAuth.GenerateAuthURL()
		if err != nil {
			return &cserrors.Error{Code: cserrors.INTERNAL_SERVER_ERROR, Message: "Error generating auth url"}
		}

		return c.Redirect(url)
	},
	)

	authRouter.Get("/sign-in/google/callback", func(c *fiber.Ctx) error {
		state := c.Query("state")
		if !googleAuth.VerifyState(state) {
			return &cserrors.Error{Code: cserrors.BAD_REQUEST, Message: "Invalid State"}
		}

		ctx := context.Background()

		code := c.Query("code")

		userInfo, err := googleAuth.GetUserInfo(ctx, code)
		if err != nil {
			return &cserrors.Error{Code: cserrors.INTERNAL_SERVER_ERROR, Message: "Error getting user info"}
		}

		user, err := userService.GetByEmail(c.Context(), userInfo.Email)
		if err != nil {
			return &cserrors.Error{Code: cserrors.UNAUTHORIZED, Message: "Unauthorized"}
		}

		token, err := auth.SignJWT(user, appConfig.JWTSecret)
		if err != nil {
			return &cserrors.Error{Code: cserrors.INTERNAL_SERVER_ERROR, Message: "Error signing JWT"}
		}

		return c.JSON(fiber.Map{
			"message": "OK",
			"token":   token,
		})
	})

	authRouter.Post("/sign-in/credential", middlewares.ValidateMiddleware(&requests.Credential{}), func(c *fiber.Ctx) error {
		credential := c.Locals("request").(*requests.Credential)

		user, err := userService.GetByUsername(c.Context(), credential.Username)
		if err != nil {
			return &cserrors.Error{Code: cserrors.UNAUTHORIZED, Message: "Unauthorized"}
		}

		password, err := userService.GetPasswordByID(c.Context(), user.ID)
		if err != nil {
			return &cserrors.Error{Code: cserrors.UNAUTHORIZED, Message: "Unauthorized"}
		}

		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(credential.Password))
		if err != nil {
			return &cserrors.Error{Code: cserrors.UNAUTHORIZED, Message: "Unauthorized"}
		}

		token, err := auth.SignJWT(user, appConfig.JWTSecret)
		if err != nil {
			return &cserrors.Error{Code: cserrors.UNAUTHORIZED, Message: "Unauthorized"}
		}

		return c.JSON(fiber.Map{
			"message": "OK",
			"token":   token,
		})
	})

}
