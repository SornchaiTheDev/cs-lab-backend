package services

import (
	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
)

type UserService interface {
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetPasswordByID(ID string) (string, error)
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
