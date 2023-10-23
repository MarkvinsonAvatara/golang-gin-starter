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
	HandledBy   sql.NullString `json:"handled_by"`
	RequestedAt time.Time      `json:"requested_at"`
	HandledAt   time.Time      `json:"handled_at"`
}

func NewUditablePinjaman(requestedBy string) AuditablePinjaman {
	return AuditablePinjaman{
		RequestedAt: time.Now(),
		RequestedBy: utils.StringToNullString(requestedBy),
	}
}

func NewAuditableBook(createdBy string) Auditable {
	return Auditable{
		CreatedAt: time.Now(),
		CreatedBy: utils.StringToNullString(createdBy),
	}
}

func NewUpdateAuditableBook(updatedBy string) Auditable {
	return Auditable{
		UpdatedAt: time.Now(),
		UpdatedBy: utils.StringToNullString(updatedBy),
	}
}

func NewHandledPinjaman(handledBy string) AuditablePinjaman {
	return AuditablePinjaman{
		HandledAt: time.Now(),
		HandledBy: utils.StringToNullString(handledBy),
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

func NewUserAuditable(createdBy string) Auditable {
	return Auditable{
		CreatedAt: time.Now(),
		CreatedBy: utils.StringToNullString(createdBy),
	}
}

func NewUserUpdateAuditable(updatedBy string) Auditable {
	return Auditable{
		UpdatedAt: time.Now(),
		UpdatedBy: utils.StringToNullString(updatedBy),
	}
}


func NewAuditableUserRole(createdBy string) Auditable {
	return Auditable{
		CreatedAt: time.Now(),
		CreatedBy: utils.StringToNullString(createdBy),
	}
}

func NewAuditableUserRoleUpdate(updatedBy string) Auditable {
	return Auditable{
		UpdatedAt: time.Now(),
		UpdatedBy: utils.StringToNullString(updatedBy),
	}
}
