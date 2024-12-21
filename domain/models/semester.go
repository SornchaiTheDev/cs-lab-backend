package models

import (
	"time"

	"github.com/SornchaiTheDev/cs-lab-backend/constants"
)

type Semester struct {
	ID        string                 `json:"id" db:"id"`
	Name      string                 `json:"name" db:"name"`
	Type      constants.SemesterType `json:"type" db:"type"`
	StartDate time.Time              `json:"started_date" db:"started_date"`
	CreatedAt time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt time.Time              `json:"updated_at" db:"updated_at"`
	IsDeleted bool                   `json:"is_deleted" db:"is_deleted"`
	DeletedAt *time.Time             `json:"deleted_at" db:"deleted_at"`
}
