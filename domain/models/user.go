package models

import "time"

type User struct {
	Username     string    `db:"username"`
	DisplayName  string    `db:"display_name"`
	ProfileImage string    `db:"profile_image"`
	Roles        []string  `db:"roles"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	DeleteAt     time.Time `db:"delete_at"`
}
