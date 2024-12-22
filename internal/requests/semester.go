package requests

import "time"

type Semester struct {
	Name        string    `json:"name" db:"name" validate:"required,min=1"`
	Type        string    `json:"type" db:"type" validate:"required,oneof=first second summer"`
	StartedDate time.Time `json:"started_date" db:"started_date" validate:"required"`
}
