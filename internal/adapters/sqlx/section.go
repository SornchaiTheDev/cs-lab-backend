package sqlx

import (
	"context"
	"errors"
	"fmt"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type sectionRepository struct {
	db *sqlx.DB
}

func NewSqlxSectionRepository(db *sqlx.DB) repositories.SectionRepository {
	return &sectionRepository{db}
}

func (r *sectionRepository) Create(ctx context.Context, sec *requests.Section) (*models.Section, error) {
	query := `INSERT INTO sections (name, started_at, ended_at, icon, course_id, semester_id)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`

	row := r.db.QueryRowxContext(ctx, query, sec.Name, sec.StartedAt, sec.EndedAt, sec.Icon, sec.CourseID, sec.SemesterID)

	var section models.Section
	err := row.StructScan(&section)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return nil, cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Section already exists")
			}
		}
		return nil, err
	}

	return &section, nil
}

func (r *sectionRepository) GetPagination(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]models.Section, error) {
	query := fmt.Sprintf(`SELECT * FROM sections 
		WHERE (name LIKE $1)
		AND deleted_at IS NULL
		ORDER BY %s %s
		OFFSET $2
		LIMIT $3
		`, sortBy, sortOrder)

	rows, err := r.db.QueryxContext(ctx, query, "%"+search+"%", (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}

	sections := []models.Section{}

	for rows.Next() {
		var sem models.Section
		err = rows.StructScan(&sections)
		if err != nil {
			return nil, err
		}

		sections = append(sections, sem)
	}

	return sections, nil
}

func (r *sectionRepository) Count(ctx context.Context, search string) (int, error) {
	query := `
		SELECT COUNT(*) FROM sections 
		WHERE (name LIKE $1)
		AND deleted_at IS NULL
	`
	row := r.db.QueryRowContext(ctx, query, "%"+search+"%")

	var count int

	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *sectionRepository) GetByID(ctx context.Context, ID string) (*models.Section, error) {
	return nil, nil
}

func (r *sectionRepository) UpdateByID(ctx context.Context, ID string, sec *requests.Section) (*models.Section, error) {
	return nil, nil
}

func (r *sectionRepository) DeleteByID(ctx context.Context, ID string) error {
	return nil
}
