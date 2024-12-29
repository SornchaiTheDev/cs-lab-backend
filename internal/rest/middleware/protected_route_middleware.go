package middleware

import (
	"fmt"

	"github.com/SornchaiTheDev/cs-lab-backend/infrastructure/auth"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/rerror"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedRouteMiddleware(secret string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access_token")

		token, err := jwt.ParseWithClaims(accessToken, &auth.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(secret), nil
		})

		if err != nil {
			return rerror.Res(c, rerror.ERR_UNAUTHORIZED, "Unauthorized")
		}

		if claims, ok := token.Claims.(*auth.JWTClaims); ok {
			c.Locals("user", claims)
			return c.Next()
		}

		return rerror.Res(c, rerror.ERR_UNAUTHORIZED, "Unauthorized")
	}
}
