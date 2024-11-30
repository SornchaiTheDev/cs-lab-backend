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
