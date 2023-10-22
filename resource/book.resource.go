package resource

import (
	"gin-starter/entity"
)

type BookDetail struct {
	ID          string `json:"id"`
	Isbn        int64  `json:"isbn"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Edition     int64  `json:"edition"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetBookByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetBookRequest struct {
	Search string `form:"search" json:"search"`
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Limit  int    `form:"limit,default=10" json:"limit"`
	Page   int    `form:"page,default=0" json:"page"`
}

type CreateBookRequest struct {
	Isbn        int64  `form:"isbn" json:"isbn"`
	Title       string `form:"title" json:"title"`
	Genre       string `form:"genre" json:"genre"`
	Author      string `form:"author" json:"author"`
	Publisher   string `form:"publisher" json:"publisher"`
	Edition     int64  `form:"edition" json:"edition"`
	Description string `form:"description" json:"description"`
}

type BookResponse struct {
	ID        string `json:"id"`
	Isbn      int64  `json:"isbn"`
	Title     string `json:"title"`
	Genre     string `json:"genre"`
	Author    string `json:"author"`
	Desc      string `json:"desc"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateBookRequest struct {
	ID          string `form:"id" json:"id"`
	Isbn        int64  `form:"isbn" json:"isbn"`
	Title       string `form:"title" json:"title"`
	Genre       string `form:"genre" json:"genre"`
	Author      string `form:"author" json:"author"`
	Publisher   string `form:"publisher" json:"publisher"`
	Edition     int64  `form:"edition" json:"edition"`
	Description string `form:"description" json:"description"`
}

type GetBookListResponse struct {
	List []*BookDetail `json:"list"`
	Meta *Meta         `json:"meta"`
}

type DeleteBookRequest struct {
	ID string `uri:"id" binding:"required"`
}

func NewBookResponse(book *entity.Book) *BookDetail {
	return &BookDetail{
		ID:          book.ID.String(),
		Isbn:        book.ISBN,
		Title:       book.Title,
		Genre:       book.Genre,
		Author:      book.Author,
		Publisher:   book.Publisher,
		Edition:     book.Edition,
		Description: book.Description,
		CreatedAt:   book.CreatedAt.Format(timeFormat),
		UpdatedAt:   book.UpdatedAt.Format(timeFormat),
	}
}
