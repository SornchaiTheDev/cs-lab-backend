package services

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
)

type sectionService struct {
	repo repositories.SectionRepository
}

type SectionService interface {
	Create(ctx context.Context, sem *requests.Section) (*models.Section, error)
	GetPagination(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]models.Section, error)
	Count(ctx context.Context, search string) (int, error)
	GetByID(ctx context.Context, ID string) (*models.Section, error)
	UpdateByID(ctx context.Context, ID string, sem *requests.Section) (*models.Section, error)
	DeleteByID(ctx context.Context, ID string) error
}

func NewSectionService(repo repositories.SectionRepository) SectionService {
	return &sectionService{
		repo: repo,
	}
}

func (s *sectionService) Create(ctx context.Context, sem *requests.Section) (*models.Section, error) {
	return s.repo.Create(ctx, sem)
}

func (s *sectionService) GetPagination(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]models.Section, error) {
	sanitizedSortBy, err := sanitizeSortBy(sortBy, &models.Section{})
	if err != nil {
		return nil, err
	}

	sanitizedSortOrder := sanitizeSortOrder(sortOrder)

	return s.repo.GetPagination(ctx, page, limit, search, sanitizedSortBy, sanitizedSortOrder)

}

func (s *sectionService) Count(ctx context.Context, search string) (int, error) {
	return s.repo.Count(ctx, search)
}

func (s *sectionService) GetByID(ctx context.Context, ID string) (*models.Section, error) {
	return s.repo.GetByID(ctx, ID)
}

func (s *sectionService) UpdateByID(ctx context.Context, ID string, sem *requests.Section) (*models.Section, error) {
	return s.repo.UpdateByID(ctx, ID, sem)
}

func (s *sectionService) DeleteByID(ctx context.Context, ID string) error {
	return s.repo.DeleteByID(ctx, ID)
}
