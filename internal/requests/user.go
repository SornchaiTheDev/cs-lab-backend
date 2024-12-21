package requests

type User struct {
	Username    string   `json:"username" db:"username"`
	Email       string   `json:"email" db:"email"`
	DisplayName string   `json:"display_name" db:"display_name"`
	Roles       []string `json:"roles" db:"roles"`
}

type CredentialUser struct {
	User
	Password string `json:"password"`
}
