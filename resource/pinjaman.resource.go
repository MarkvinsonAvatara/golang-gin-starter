package resource

import (
	"gin-starter/entity"
)

type PinjamanDetail struct {
	ID          string `json:"id"`
	UserId      string `json:"userid"`
	BukuId      string `json:"bukuid"`
	TglPinjam   string `json:"tglpinjam"`
	TglKembali  string `json:"tglkembali"`
	Status      bool   `json:"status"`
	RequestedAt string `json:"requested_at"`
	HandledAt   string `json:"handled_at"`
}

type GetPinjamanByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type CreatePinjamanRequest struct {
	UserId     string `form:"userid" json:"userid"`
	BukuId     string `form:"bukuid" json:"bukuid"`
	TglPinjam  string `form:"tglpinjam" json:"tglpinjam"`
	TglKembali string `form:"tglkembali" json:"tglkembali"`
}

type GetPinjamanListResponse struct {
	List  []*PinjamanDetail `json:"list"`
	Total int64             `json:"total"`
}

type GetPinjamanRequest struct {
	Query string `form:"query" json:"query"`
	Sort  string `form:"sort" json:"sort"`
	Order string `form:"order" json:"order"`
	Limit int    `form:"limit,default=10" json:"limit"`
	Page  int    `form:"page,default=0" json:"page"`
}

type PinjamanRespons struct {
	ID          string `json:"id"`
	UserId      string `json:"userid"`
	BukuId      string `json:"bukuid"`
	TglPinjam   string `json:"tglpinjam"`
	TglKembali  string `json:"tglkembali"`
	Status      bool   `json:"status"`
	RequestedAt string `json:"requested_at"`
	HandledAt   string `json:"handled_at"`
}

type HandledRequest struct {
	ID     string `form:"id" json:"id"`
	Status bool   `form:"status" json:"status"`
}

func NewPinjamanResponse(pinjaman *entity.Pinjaman) *PinjamanDetail {
	return &PinjamanDetail{
		ID:          pinjaman.ID.String(),
		UserId:      pinjaman.UserID,
		BukuId:      pinjaman.BukuID,
		TglPinjam:   pinjaman.Tglpinjam.Time.Format(timeFormat),
		TglKembali:  pinjaman.Tglkembali.Time.Format(timeFormat),
		Status:      pinjaman.Status,
		RequestedAt: pinjaman.RequestedAt.Format(timeFormat),
		HandledAt:   pinjaman.HandledAt.Format(timeFormat),
	}
}
