package repositories

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
)

type SemesterRepository interface {
	Create(ctx context.Context, sem *requests.Semester) (*models.Semester, error)
}
