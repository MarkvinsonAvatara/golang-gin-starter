package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"

	"gin-starter/utils"
)

// Auditable is an interface that can be embedded in structs that need to be auditable
type Auditable struct {
	CreatedBy sql.NullString `json:"created_by"`
	UpdatedBy sql.NullString `json:"updated_by"`
	DeletedBy sql.NullString `json:"deleted_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type AuditablePinjaman struct {
	RequestedBy sql.NullString `json:"requested_by"`
	HandledBy  sql.NullString `json:"handled_by"`
	RequestedAt time.Time      `json:"requested_at"`
	HandledAt   time.Time      `json:"handled_at"`
}

func NewUditablePinjaman(requestedBy string) AuditablePinjaman {
	return AuditablePinjaman{
		RequestedAt: time.Now(),
		HandledAt:   time.Now(),
		RequestedBy: utils.StringToNullString(requestedBy),
		HandledBy:   utils.StringToNullString(requestedBy),
	}
}

// NewAuditable creates a new Auditable struct
func NewAuditable(createdBy string) Auditable {
	return Auditable{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: utils.StringToNullString(createdBy),
		UpdatedBy: utils.StringToNullString(createdBy),
	}
}
