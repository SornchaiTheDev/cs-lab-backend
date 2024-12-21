package sqlx

import (
	"context"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/jmoiron/sqlx"
)

type sqlxSemesterRepository struct {
	db *sqlx.DB
}

func NewSqlxSemesterRepository(db *sqlx.DB) repositories.SemesterRepository {
	return &sqlxSemesterRepository{
		db: db,
	}
}

func (r *sqlxSemesterRepository) Create(ctx context.Context, sem *requests.Semester) (*models.Semester, error) {
	query := `INSERT INTO semesters (
		name,
		type,
		started_date
	) VALUES ($1,$2,$3)
	RETURNING *`
	row := r.db.QueryRowxContext(ctx, query, sem.Name, sem.Type, sem.StartedDate)

	var semester models.Semester
	err := row.StructScan(&semester)
	if err != nil {
		return nil, err
	}

	return &semester, nil

}
