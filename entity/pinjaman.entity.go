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
	Userid     *User        `gorm:"foreignKey:id" json:"user"`
	Bookid     *Book        `gorm:"foreignKey:id" json:"book"`
	Tglpinjam  sql.NullTime `json:"tglpinjam"`
	Tglkembali sql.NullTime `json:"tglkembali"`
	Status     string       `json:"status"`
	Auditable
}

// TableName specifies table name
func (model *Pinjaman) TableName() string {
	return pinjamanTableNm
}

// NewPinjaman create new entity Pinjaman
func NewPinjaman(
	id uuid.UUID,
	tglpinjam sql.NullTime,
	tglkembali sql.NullTime,
	status string,
	createdBy string,
) *Pinjaman {
	return &Pinjaman{
		ID:         id,
		Tglpinjam:  tglpinjam,
		Tglkembali: tglkembali,
		Status:     status,
		Auditable:  NewAuditable(createdBy),
	}
}

func UpdatePinjaman(
	id uuid.UUID,
	tglkembali sql.NullTime,
	status string,
	updatedBy string,
) *Pinjaman {
	return &Pinjaman{
		ID:         id,
		Tglkembali: tglkembali,
		Status:     status,
		Auditable:  NewAuditable(updatedBy),
	}
}

func (model *Pinjaman) MapUpdateFrom(from *Pinjaman) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"tglkembali": model.Tglkembali,
			"status":     model.Status,
			"updated_at": model.UpdatedAt,
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
