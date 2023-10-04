package entity

import (
	"github.com/google/uuid"
)

const (
	bukuTableName = "public.buku"
)

type Buku struct {
	ID          uuid.UUID `json:"id"`
	ISBN        int64     `json:"isbn"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	Edition     int64     `json:"edition"`
	Year        int64     `json:"year"`
	Description string    `json:"description"`
}

func (u *Buku) TableName() string {
	return bukuTableName
}
