package services

import (
	"context"
	"fmt"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetPasswordByID(ID string) (string, error)
	GetPagination(page int, limit int, search string) ([]models.User, error)
	Count() (int, error)
	Create(ctx context.Context, user *requests.CreateUser) (*models.User, error)
	SetPassword(ctx context.Context, username string, password string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetByEmail(email string) (*models.User, error) {
	return s.userRepository.GetByEmail(email)
}

func (s *userService) GetByUsername(username string) (*models.User, error) {
	return s.userRepository.GetByUsername(username)
}

func (s *userService) GetPasswordByID(ID string) (string, error) {
	return s.userRepository.GetPasswordByID(ID)
}

func (s *userService) GetPagination(page int, limit int, search string) ([]models.User, error) {
	return s.userRepository.GetPagination(page, limit, search)
}

func (s *userService) Count() (int, error) {
	return s.userRepository.Count()
}

func (s *userService) Create(ctx context.Context, user *requests.CreateUser) (*models.User, error) {
	return s.userRepository.Create(ctx, user)
}

func (s *userService) SetPassword(ctx context.Context, username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return fmt.Errorf("Cannot generate password")
	}

	return s.userRepository.SetPassword(ctx, username, string(hashedPassword))
}
