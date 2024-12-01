package repositories

import "github.com/SornchaiTheDev/cs-lab-backend/domain/models"

type UserRepository interface {
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetPasswordByID(ID string) (string, error)
}
