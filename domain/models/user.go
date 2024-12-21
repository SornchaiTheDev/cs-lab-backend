package models

import (
	"time"
)

type User struct {
	ID           string     `db:"id" json:"id"`
	Username     string     `db:"username" json:"username"`
	Email        string     `db:"email" json:"email"`
	DisplayName  string     `db:"display_name" json:"display_name"`
	ProfileImage *string    `db:"profile_image" json:"profile_image"`
	Roles        []string   `db:"roles" json:"roles"`
	IsDeleted    bool       `db:"is_deleted" json:"is_deleted"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at"`
}
