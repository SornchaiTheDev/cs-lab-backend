package services

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
)

type SemesterService interface {
	Create(ctx context.Context, sem *requests.Semester) (*models.Semester, error)
	GetPagination(ctx context.Context, page int, limit int, search string) ([]models.Semester, error)
	Count(ctx context.Context, search string) (int, error)
}

type semesterService struct {
	repo repositories.SemesterRepository
}

func NewSemesterService(repo repositories.SemesterRepository) *semesterService {
	return &semesterService{
		repo: repo,
	}
}

func (s *semesterService) Create(ctx context.Context, sem *requests.Semester) (*models.Semester, error) {
	return s.repo.Create(ctx, sem)
}

func (s *semesterService) GetPagination(ctx context.Context, page int, limit int, search string) ([]models.Semester, error) {
	return s.repo.GetPagination(ctx, page, limit, search)
}

func (s *semesterService) Count(ctx context.Context, search string) (int, error) {
	return s.repo.Count(ctx, search)
}
