package resource

import (
	"gin-starter/entity"
)

type PinjamanDetail struct {
	ID          string `json:"id"`
	UserId      string `json:"userid"`
	Name        string `json:"name"`
	DOB         string `json:"dob"`
	BukuId      string `json:"bukuid"`
	ISBN        int64  `json:"isbn"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Edition     int64  `json:"edition"`
	// Description string `json:"description"`
	TglPinjam   string `json:"tglpinjam"`
	TglKembali  string `json:"tglkembali"`
	Status      int64  `json:"status"`
	RequestedAt string `json:"requested_at"`
	HandledAt   string `json:"handled_at"`
}

type GetPinjamanByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type CreatePinjamanRequest struct {
	UserId     string `form:"userid" json:"userid" `
	BukuId     string `form:"bukuid" json:"bukuid" `
	TglPinjam  string `form:"tglpinjam" json:"tglpinjam" `
	TglKembali string `form:"tglkembali" json:"tglkembali" `
}

type GetPinjamanListResponse struct {
	List []*PinjamanDetail `json:"list"`
	Meta *Meta             `json:"meta"`
}

type GetPinjamanRequest struct {
	Search string `form:"search" json:"search"`
	Filter string `form:"filter" json:"filter"`
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Limit  int    `form:"limit,default=10" json:"limit"`
	Page   int    `form:"page,default=0" json:"page"`
}

type PinjamanRespons struct {
	ID          string `json:"id"`
	UserId      string `json:"userid"`
	BukuId      string `json:"bukuid"`
	TglPinjam   string `json:"tglpinjam"`
	TglKembali  string `json:"tglkembali"`
	Status      int64  `json:"status"`
	RequestedAt string `json:"requested_at"`
	HandledAt   string `json:"handled_at"`
}

type HandledRequest struct {
	ID     string `form:"id" json:"id"`
	Status int64  `form:"status" json:"status" binding:"required"`
}

func NewPinjamanResponse(pinjaman *entity.PinjamanDetail) *PinjamanDetail {
	return &PinjamanDetail{
		ID:          pinjaman.ID.String(),
		UserId:      pinjaman.UserID,
		Name:        pinjaman.Name,
		DOB:         pinjaman.DOB.Time.Format(timeFormat),
		BukuId:      pinjaman.BukuID,
		ISBN:        pinjaman.ISBN,
		Title:       pinjaman.Title,
		Genre:       pinjaman.Genre,
		Author:      pinjaman.Author,
		Publisher:   pinjaman.Publisher,
		Edition:     pinjaman.Edition,
		// Description: pinjaman.BukuDetail.Description,
		TglPinjam:   pinjaman.Tglpinjam.Time.Format(timeFormat),
		TglKembali:  pinjaman.Tglkembali.Time.Format(timeFormat),
		Status:      pinjaman.Status,
		RequestedAt: pinjaman.RequestedAt.Format(timeFormat),
		HandledAt:   pinjaman.HandledAt.Format(timeFormat),
	}
}

func CreateNewPinjamanResponse(pinjaman *entity.Pinjaman) *PinjamanRespons {
	return &PinjamanRespons{
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
