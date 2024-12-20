package auth

import (
	"time"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Sub          string        `json:"sub"`
	DisplayName  string        `json:"displayName"`
	ProfileImage string        `json:"profileImage"`
	Roles        []interface{} `json:"roles"`
	Iss          string        `json:"iss"`
	Exp          int64         `json:"exp"`
}

func SignJWT(user *models.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":          user.ID,
		"displayName":  user.DisplayName,
		"profileImage": user.ProfileImage,
		"roles":        user.Roles,
		"iss":          "cs-lab-backend",
		"exp":          time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(secret))

}
