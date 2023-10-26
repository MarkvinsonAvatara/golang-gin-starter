package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const (
	pinjamanTableNm = "public.pinjaman"
)

// Pinjaman define for table pinjaman

type Pinjaman struct {
	ID         uuid.UUID    `json:"id"`
	UserID     string       `gorm:"foreignKey:id" json:"userid"`
	BukuID     string       `gorm:"foreignKey:id" json:"bukuid"`
	Tglpinjam  sql.NullTime `json:"tglpinjam"`
	Tglkembali sql.NullTime `json:"tglkembali"`
	Status     int64        `json:"status"`
	AuditablePinjaman
}

// TableName specifies table name
func (model *Pinjaman) TableName() string {
	return pinjamanTableNm
}

// NewPinjaman create new entity Pinjaman
func NewPinjaman(
	id uuid.UUID,
	userid string,
	bukuid string,
	tglpinjam sql.NullTime,
	tglkembali sql.NullTime,
	requestedBy string,
) *Pinjaman {
	return &Pinjaman{
		ID:                id,
		UserID:            userid,
		BukuID:            bukuid,
		Tglpinjam:         tglpinjam,
		Tglkembali:        tglkembali,
		AuditablePinjaman: NewUditablePinjaman(requestedBy),
	}
}

func HandledPinjaman(
	id uuid.UUID,
	handledBy string,
	status int64,
) *Pinjaman {
	return &Pinjaman{
		ID:                id,
		AuditablePinjaman: NewHandledPinjaman(handledBy),
		Status:            status,
	}
}

func (model *Pinjaman) MapUpdateFrom(from *Pinjaman) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"tglkembali": model.Tglkembali,
			"status":     model.Status,
			"handled_at": model.HandledAt,
			"handled_by": model.HandledBy,
		}
	}

	mapped := make(map[string]interface{})

	if model.Tglkembali != from.Tglkembali {
		mapped["tglkembali"] = model.Tglkembali
	}

	if model.Status != from.Status {
		mapped["status"] = model.Status
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
