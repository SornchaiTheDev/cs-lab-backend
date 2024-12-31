package models

type Course struct {
	ID        string `json:"id" db:"id"`
	Code      string `json:"code" db:"code"`
	Name      string `json:"name" db:"name"`
	CreatedBy string `json:"created_by" db:"created_by"`
	RecordStatus
}
