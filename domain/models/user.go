package models

type User struct {
	ID           string   `db:"id" json:"id"`
	Username     string   `db:"username" json:"username"`
	Email        string   `db:"email" json:"email"`
	DisplayName  string   `db:"display_name" json:"display_name"`
	ProfileImage *string  `db:"profile_image" json:"profile_image"`
	Roles        []string `db:"roles" json:"roles"`
	RecordStatus
}
