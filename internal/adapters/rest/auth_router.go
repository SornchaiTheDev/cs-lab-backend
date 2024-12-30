package rest

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/infrastructure/auth"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/middlewares"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// TODO: redirect back to frontend and set cookie
func NewAuthRouter(router fiber.Router, appConfig *configs.Config, userService services.UserService, refreshTokenService services.RefreshTokenService) {
	authRouter := router.Group("/auth")

	googleAuth := auth.NewGoogleAuth(appConfig)

	// Google OAuth2
	authRouter.Get("/sign-in/google", func(c *fiber.Ctx) error {
		url, err := googleAuth.GenerateAuthURL()
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error generating auth url")
		}

		return c.Redirect(url)
	})

	authRouter.Get("/sign-in/google/callback", func(c *fiber.Ctx) error {
		state := c.Query("state")
		if !googleAuth.VerifyState(state) {
			return cserrors.New(cserrors.BAD_REQUEST, "Invalid State")
		}

		ctx := context.Background()

		code := c.Query("code")

		userInfo, err := googleAuth.GetUserInfo(ctx, code)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error getting user info")
		}

		user, err := userService.GetByEmail(c.Context(), userInfo.Email)
		if err != nil {
			return cserrors.New(cserrors.UNAUTHORIZED, "Unauthorized")
		}

		newAccessToken, err := auth.SignAccessToken(user, appConfig.JWTSecret)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		newRefreshToken, err := auth.SignRefreshToken(user.ID, appConfig.JWTRefreshSecret)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		err = refreshTokenService.Set(c.Context(), user.ID, newRefreshToken)
		if err != nil {

			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		return c.JSON(fiber.Map{
			"message":       "OK",
			"access_token":  newAccessToken,
			"refresh_token": newRefreshToken,
		})
	})

	authRouter.Post("/sign-in/credential", middlewares.ValidateMiddleware(&requests.Credential{}), func(c *fiber.Ctx) error {
		credential := c.Locals("request").(*requests.Credential)

		user, err := userService.GetByUsername(c.Context(), credential.Username)
		if err != nil {
			return cserrors.New(cserrors.UNAUTHORIZED, "Unauthorized")
		}

		password, err := userService.GetPasswordByID(c.Context(), user.ID)
		if err != nil {
			return cserrors.New(cserrors.UNAUTHORIZED, "Unauthorized")
		}

		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(credential.Password))
		if err != nil {
			return cserrors.New(cserrors.UNAUTHORIZED, "Unauthorized")
		}

		newAccessToken, err := auth.SignAccessToken(user, appConfig.JWTSecret)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		newRefreshToken, err := auth.SignRefreshToken(user.ID, appConfig.JWTRefreshSecret)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		err = refreshTokenService.Set(c.Context(), user.ID, newRefreshToken)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		return c.JSON(fiber.Map{
			"message":       "OK",
			"access_token":  newAccessToken,
			"refresh_token": newRefreshToken,
		})
	})

	authRouter.Post("/refresh-token", middlewares.ProtectedRouteMiddleware(appConfig.JWTSecret), middlewares.ValidateMiddleware(&requests.RefreshToken{}), func(c *fiber.Ctx) error {
		request := c.Locals("request").(*requests.RefreshToken)
		user := c.Locals("user").(*models.User)

		refreshToken, err := refreshTokenService.Get(c.Context(), user.ID)
		if err != nil {
			return cserrors.New(cserrors.UNAUTHORIZED, "Unauthorized")
		}

		// check for replay attack
		if refreshToken != request.RefreshToken {
			return cserrors.New(cserrors.UNAUTHORIZED, "Unauthorized")
		}

		// verify refresh token if valid and not expired
		err = auth.VerifyAccessToken(request.RefreshToken, appConfig.JWTRefreshSecret)
		if err != nil {
			return cserrors.New(cserrors.UNAUTHORIZED, "Unauthorized")
		}

		newAccessToken, err := auth.SignAccessToken(user, appConfig.JWTSecret)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		newRefreshToken, err := auth.SignRefreshToken(user.ID, appConfig.JWTRefreshSecret)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		err = refreshTokenService.Set(c.Context(), user.ID, newRefreshToken)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Something went wrong")
		}

		return c.JSON(fiber.Map{
			"message":       "OK",
			"access_token":  newAccessToken,
			"refresh_token": newRefreshToken,
		})
	})

}
