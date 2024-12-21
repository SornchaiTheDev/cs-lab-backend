package requests

import "time"

type Semester struct {
	Name        string    `json:"name" validate:"required,min=1"`
	Type        string    `json:"type" validate:"required,oneof=first second summer"`
	StartedDate time.Time `json:"started_date" validate:"required"`
}
