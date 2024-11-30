package services

import (
	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
)

type UserService interface {
	GetUserByEmail(email string) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetByEmail(email)
}
