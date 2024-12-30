package services

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
)

type RefreshTokenService interface {
	Get(ctx context.Context, userID string) (string, error)
	Set(ctx context.Context, userID string, token string) error
}

type refreshTokenService struct {
	repo repositories.RefreshTokenRepository
}

func NewRefreshTokenService(repo repositories.RefreshTokenRepository) RefreshTokenService {
	return &refreshTokenService{repo: repo}
}

func (s *refreshTokenService) Get(ctx context.Context, userID string) (string, error) {
	return s.repo.Get(ctx, userID)
}

func (s *refreshTokenService) Set(ctx context.Context, userID string, token string) error {
	return s.repo.Set(ctx, userID, token)
}
