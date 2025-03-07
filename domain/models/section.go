package models

type Section struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	StartedAt  string `json:"started_at" db:"started_at"`
	EndedAt    string `json:"ended_at" db:"ended_at"`
	Icon       string `json:"icon" db:"icon"`
	CourseID   string `json:"course_id" db:"course_id"`
	SemesterID string `json:"semester_id" db:"semester_id"`
	RecordStatus
}
