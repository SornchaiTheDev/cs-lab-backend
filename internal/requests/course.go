package requests

type Course struct {
	Name string `json:"name" db:"name" validate:"required"`
	Code string `json:"code" db:"code" validate:"required"`
}
