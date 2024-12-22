package models

type Course struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedBy string `json:"created_by" db:"created_by"`
	RecordStatus
}
