package requests

import "time"

type Section struct {
	Name       string    `json:"name" db:"name"`
	StartedAt  time.Time `json:"started_at" db:"started_at"`
	EndedAt    time.Time `json:"ended_at" db:"ended_at"`
	Icon       string    `json:"icon" db:"icon"`
	CourseID   string    `json:"course_id" db:"course_id"`
	SemesterID string    `json:"semester_id" db:"semester_id"`
}
