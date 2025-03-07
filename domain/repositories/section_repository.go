package repositories

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
)

type SectionRepository interface {
	Create(ctx context.Context, sem *requests.Section) (*models.Section, error)
	GetPagination(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]models.Section, error)
	Count(ctx context.Context, search string) (int, error)
	GetByID(ctx context.Context, ID string) (*models.Section, error)
	UpdateByID(ctx context.Context, ID string, sem *requests.Section) (*models.Section, error)
	DeleteByID(ctx context.Context, ID string) error
}
