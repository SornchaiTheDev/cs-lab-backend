package requests

type CreateUser struct {
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	DisplayName string   `json:"display_name"`
	Roles       []string `json:"roles"`
}

type CreateCredentialUser struct {
	CreateUser
	Password string `json:"password"`
}
