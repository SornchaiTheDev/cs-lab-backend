package middleware

import (
	"fmt"

	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/rerror"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedRouteMiddleware(secret string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access_token")

		token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(secret), nil
		})

		if err != nil {
			return rerror.ERR_UNAUTHORIZED
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["roles"].([]interface{}))
			// user := &auth.JWTClaims{
			// 	Sub:          claims["sub"].(string),
			// 	DisplayName:  claims["displayName"].(string),
			// 	ProfileImage: claims["profileImage"].(string),
			// 	Iss:          claims["iss"].(string),
			// 	Roles:        claims["roles"].([]string),
			// 	Exp:          int64(claims["exp"].(float64)),
			// }
			//
			// c.Locals("user", user)
		}

		return c.Next()
	}
}
