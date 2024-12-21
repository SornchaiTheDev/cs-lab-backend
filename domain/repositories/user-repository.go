package repositories

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
)

type UserRepository interface {
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetPasswordByID(ID string) (string, error)
	GetPagination(page int, limit int, search string) ([]models.User, error)
	Count() (int, error)
	Create(c context.Context, user *requests.CreateUser) (*models.User, error)
	// SetPassword(c context.Context, password string) (bool, error)
}
