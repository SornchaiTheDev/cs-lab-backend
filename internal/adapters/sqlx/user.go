package sqlx

import (
	"context"
	"fmt"
	"strings"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/models"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/repositories"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
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

func (r *sqlxUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user PostgresUser
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE email = $1", email)
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
		RecordStatus: models.RecordStatus{
			IsDeleted: user.IsDeleted,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
	}, nil
}

func (r *sqlxUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user PostgresUser
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1", username)
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
		RecordStatus: models.RecordStatus{
			IsDeleted: user.IsDeleted,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
	}, nil
}

func (r *sqlxUserRepository) GetPasswordByID(ctx context.Context, ID string) (string, error) {
	row := r.db.QueryRowContext(ctx, "SELECT password FROM user_passwords WHERE user_id = $1", ID)
	var password string

	err := row.Scan(&password)

	if err != nil {
		return "", err
	}

	return password, nil
}

func (r *sqlxUserRepository) GetPagination(ctx context.Context, page int, limit int, search string) ([]models.User, error) {
	rows, err := r.db.QueryxContext(ctx, `SELECT * FROM users 
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
			RecordStatus: models.RecordStatus{
				IsDeleted: user.IsDeleted,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
		})
	}

	return users, nil
}

func (r *sqlxUserRepository) Count(ctx context.Context, search string) (int, error) {
	query := `
		SELECT COUNT(*) FROM users 
		WHERE (username LIKE $1 
		OR display_name LIKE $1 
		OR email LIKE $1) AND
	        deleted_at IS NULL
	`
	row := r.db.QueryRowContext(ctx, query, "%"+search+"%")

	var count int

	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *sqlxUserRepository) Create(ctx context.Context, user *requests.User) (*models.User, error) {
	createString := `
		INSERT INTO users (
			username,
			display_name,
			email,
			roles
		) VALUES ($1,$2,$3,string_to_array($4,',')::role[])
		RETURNING *
	`

	User := r.db.QueryRowxContext(ctx, createString, user.Username, user.DisplayName, user.Email, strings.Join(user.Roles, ","))

	var createdUser PostgresUser

	err := User.StructScan(&createdUser)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:           createdUser.ID,
		Email:        createdUser.Email,
		Username:     createdUser.Username,
		DisplayName:  createdUser.DisplayName,
		ProfileImage: createdUser.ProfileImage,
		Roles:        createdUser.Roles,
		RecordStatus: models.RecordStatus{
			IsDeleted: createdUser.IsDeleted,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
			DeletedAt: createdUser.DeletedAt,
		},
	}, nil
}

func (r *sqlxUserRepository) SetPassword(ctx context.Context, username string, password string) error {
	query := `
	INSERT INTO user_passwords (user_id,password)
	VALUES ($1,$2)
	ON CONFLICT (user_id) DO UPDATE
	SET password = $2
	`

	_, err := r.db.ExecContext(ctx, query, username, password)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqlxUserRepository) GetByID(ctx context.Context, ID string) (*models.User, error) {
	var user PostgresUser
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", ID)
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
		RecordStatus: models.RecordStatus{
			IsDeleted: user.IsDeleted,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
	}, nil
}

type updateUser struct {
	requests.User
	ID string `db:"id"`
}

func (r *sqlxUserRepository) Update(ctx context.Context, ID string, user *requests.User) (*models.User, error) {
	updateFields, err := getUpdateFields(user)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	UPDATE users
	SET %s , updated_at = NOW()
	WHERE id = :id
	RETURNING *
	`, updateFields)

	row, err := r.db.NamedQueryContext(ctx, query, &updateUser{
		ID: ID,
		User: requests.User{
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			Roles:       user.Roles,
		},
	})
	if err != nil {
		return nil, err
	}

	var updatedUser PostgresUser
	for row.Next() {
		err = row.StructScan(&updatedUser)
		if err != nil {
			return nil, err
		}
	}

	return &models.User{
		ID:           updatedUser.ID,
		Email:        updatedUser.Email,
		Username:     updatedUser.Username,
		DisplayName:  updatedUser.DisplayName,
		ProfileImage: updatedUser.ProfileImage,
		Roles:        updatedUser.Roles,
		RecordStatus: models.RecordStatus{
			IsDeleted: updatedUser.IsDeleted,
			CreatedAt: updatedUser.CreatedAt,
			UpdatedAt: updatedUser.UpdatedAt,
			DeletedAt: updatedUser.DeletedAt,
		},
	}, nil
}

func (r *sqlxUserRepository) Delete(ctx context.Context, ID string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET is_deleted = true, deleted_at = NOW() WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}
