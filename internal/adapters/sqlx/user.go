package sqlx

import (
	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type sqlxUserRepository struct {
	db *sqlx.DB
}

func NewSqlxUserRepository(db *sqlx.DB) repositories.UserRepository {
	return &sqlxUserRepository{db: db}
}

type PostgresUser struct {
	models.User
	Roles pq.StringArray `db:"roles"`
}

func (r *sqlxUserRepository) GetByEmail(email string) (*models.User, error) {
	var user PostgresUser
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		ProfileImage: user.ProfileImage,
		Roles:        user.Roles,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
	}, nil
}

func (r *sqlxUserRepository) GetByUsername(username string) (*models.User, error) {
	var user PostgresUser
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		ProfileImage: user.ProfileImage,
		Roles:        user.Roles,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
	}, nil
}

func (r *sqlxUserRepository) GetPasswordByID(ID string) (string, error) {
	row := r.db.QueryRow("SELECT password FROM user_passwords WHERE user_id = $1", ID)
	var password string

	err := row.Scan(&password)

	if err != nil {
		return "", err
	}

	return password, nil
}

func (r *sqlxUserRepository) GetPagination(page int, limit int, search string) ([]models.User, error) {
	rows, err := r.db.Queryx(`SELECT * FROM users 
		WHERE (username LIKE $1 
		OR display_name LIKE $1 
		OR email LIKE $1)
		AND deleted_at IS NULL
		OFFSET $2
		LIMIT $3
		`, "%"+search+"%", (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}

	users := []models.User{}

	for rows.Next() {
		var user PostgresUser
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, models.User{
			ID:           user.ID,
			Email:        user.Email,
			Username:     user.Username,
			DisplayName:  user.DisplayName,
			ProfileImage: user.ProfileImage,
			Roles:        user.Roles,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			DeletedAt:    user.DeletedAt,
		})
	}

	return users, nil
}

func (r *sqlxUserRepository) Count() (int, error) {
	row := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE deleted_at IS NULL")

	var count int

	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
