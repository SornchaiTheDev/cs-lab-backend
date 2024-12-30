package auth

import (
	"fmt"
	"time"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	SignAccessToken(user *models.User) (string, error)
	SignRefreshToken(userID string) (string, error)
}

type JWTClaims struct {
	DisplayName  string        `json:"displayName"`
	ProfileImage *string       `json:"profileImage"`
	Roles        []interface{} `json:"roles"`
	jwt.RegisteredClaims
}

func SignAccessToken(user *models.User, secret string) (string, error) {
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

func SignRefreshToken(userID string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": "cs-lab-backend",
		"exp": time.Now().Add(time.Hour * 24 * 5).Unix(),
	})

	return token.SignedString([]byte(secret))

}

func VerifyAccessToken(tokenString string, secret string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	return nil
}
