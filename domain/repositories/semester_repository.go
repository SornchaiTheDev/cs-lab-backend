package repositories

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
)

type SemesterRepository interface {
	Create(ctx context.Context, sem *requests.Semester) (*models.Semester, error)
	GetPagination(ctx context.Context, page int, limit int, search string) ([]models.Semester, error)
	Count(ctx context.Context, search string) (int, error)
}
