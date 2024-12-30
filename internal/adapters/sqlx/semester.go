package sqlx

import (
	"context"
	"errors"
	"fmt"

	"github.com/SornchaiTheDev/cs-lab-backend/constants"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return nil, cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Semester already exists")
			}
		}
		return nil, err
	}

	return &semester, nil

}

func (r *sqlxSemesterRepository) GetPagination(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]models.Semester, error) {
	query := fmt.Sprintf(`SELECT * FROM semesters 
		WHERE (name LIKE $1 
		OR type::text LIKE $1
		OR DATE(started_date)::text = $1)
		AND deleted_at IS NULL
		ORDER BY %s %s
		OFFSET $2
		LIMIT $3
		`, sortBy, sortOrder)

	rows, err := r.db.QueryxContext(ctx, query, "%"+search+"%", (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}

	sems := []models.Semester{}

	for rows.Next() {
		var sem models.Semester
		err = rows.StructScan(&sem)
		if err != nil {
			return nil, err
		}

		sems = append(sems, sem)
	}

	return sems, nil
}

func (r *sqlxSemesterRepository) Count(ctx context.Context, search string) (int, error) {
	query := `
		SELECT COUNT(*) FROM semesters 
		WHERE (name LIKE $1
		OR type::text LIKE $1 
		OR DATE(started_date)::text = $1) AND deleted_at IS NULL
	`
	row := r.db.QueryRowContext(ctx, query, "%"+search+"%")

	var count int

	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *sqlxSemesterRepository) GetByID(ctx context.Context, ID string) (*models.Semester, error) {
	row := r.db.QueryRowxContext(ctx, "SELECT * FROM semesters WHERE id = $1", ID)

	var sem models.Semester

	err := row.StructScan(&sem)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "22P02" {
				return nil, cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Semester not found")
			}
		}
		return nil, err
	}

	return &sem, nil
}

func (r *sqlxSemesterRepository) UpdateByID(ctx context.Context, ID string, sem *requests.Semester) (*models.Semester, error) {
	updateFields := getUpdateFields(sem)

	query := fmt.Sprintf(`
	UPDATE semesters
	SET %s , updated_at = NOW()
	WHERE id = :id
	RETURNING *
	`, updateFields)

	row, err := r.db.NamedQueryContext(ctx, query, &models.Semester{
		ID:        ID,
		Name:      sem.Name,
		StartDate: sem.StartedDate,
		Type:      constants.SemesterType(sem.Type),
	})
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "22P02" {
				return nil, cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Semester not found")
			}
		}
		return nil, err
	}

	var updatedSem models.Semester
	for row.Next() {
		err = row.StructScan(&updatedSem)
		if err != nil {
			return nil, err
		}
	}

	return &updatedSem, nil
}

func (r *sqlxSemesterRepository) DeleteByID(ctx context.Context, ID string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE semesters SET is_deleted = true, deleted_at = NOW() WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}
