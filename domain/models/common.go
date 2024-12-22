package models

import "time"

type RecordStatus struct {
	IsDeleted bool       `db:"is_deleted" json:"is_deleted"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
