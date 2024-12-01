package auth

import (
	"time"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

func SignJWT(user *models.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":           user.ID,
		"displayName":  user.DisplayName,
		"profileImage": user.ProfileImage,
		"iss":          "cs-lab-backend",
		"exp":          time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(secret))

}
