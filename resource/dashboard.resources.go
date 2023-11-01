package resource

// import (
// 	"gin-starter/entity"
// )



type GetDashboardListResponse struct {
	// List []*DashBoardDetail `json:"list"`
	Meta *DashboardMeta     `json:"meta"`
}



// type PinjamanRespons struct {
// 	ID          string `json:"id"`
// 	UserId      string `json:"userid"`
// 	BukuId      string `json:"bukuid"`
// 	TglPinjam   string `json:"tglpinjam"`
// 	TglKembali  string `json:"tglkembali"`
// 	Status      int64  `json:"status"`
// 	RequestedAt string `json:"requested_at"`
// 	HandledAt   string `json:"handled_at"`
// }

// func NewPinjamanResponse(pinjaman *entity.Pinjaman) *PinjamanDetail {
// 	return &PinjamanDetail{
// 		ID:        pinjaman.ID.String(),
// 		UserId:    pinjaman.UserID,
// 		Name:      pinjaman.Name,
// 		DOB:       pinjaman.DOB.Time.Format(timeFormat),
// 		BukuId:    pinjaman.BukuID,
// 		ISBN:      pinjaman.ISBN,
// 		Title:     pinjaman.Title,
// 		Genre:     pinjaman.Genre,
// 		Author:    pinjaman.Author,
// 		Publisher: pinjaman.Publisher,
// 		Edition:   pinjaman.Edition,
// 		// Description: pinjaman.BukuDetail.Description,
// 		TglPinjam:   pinjaman.Tglpinjam.Time.Format(timeFormat),
// 		TglKembali:  pinjaman.Tglkembali.Time.Format(timeFormat),
// 		Status:      pinjaman.Status,
// 		RequestedAt: pinjaman.RequestedAt.Format(timeFormat),
// 		HandledAt:   pinjaman.HandledAt.Format(timeFormat),
// 	}
// }
