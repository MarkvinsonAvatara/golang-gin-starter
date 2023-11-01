package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	bookTableName = "public.book"
)

type Book struct {
	ID          uuid.UUID `json:"id"`
	ISBN        int64     `json:"isbn"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	Edition     int64     `json:"edition"`
	Description string    `json:"description"`
	Auditable
	Status int64 `json:"status"`
}

type BookPinjaman struct {
	ID          uuid.UUID `json:"id"`
	ISBN        int64     `json:"isbn"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	Edition     int64     `json:"edition"`
	Description string    `json:"description"`
	Status      int64     `json:"status"`
}

func (u *Book) TableName() string {
	return bookTableName
}

func NewBook(
	id uuid.UUID,
	isbn int64,
	title string,
	author string,
	genre string,
	publisher string,
	edition int64,
	description string,
	createdBy string,
) *Book {
	return &Book{
		ID:          id,
		ISBN:        isbn,
		Title:       title,
		Author:      author,
		Genre:       genre,
		Publisher:   publisher,
		Edition:     edition,
		Description: description,
		Auditable:   NewAuditableBook(createdBy),
	}
}

func UpdateBook(
	id uuid.UUID,
	isbn int64,
	title string,
	author string,
	genre string,
	publisher string,
	edition int64,
	description string,
	updatedBy string,
) *Book {
	return &Book{
		ID:          id,
		ISBN:        isbn,
		Title:       title,
		Author:      author,
		Genre:       genre,
		Publisher:   publisher,
		Edition:     edition,
		Description: description,
		Auditable:   NewUpdateAuditableBook(updatedBy),
	}
}

// MapUpdateFrom mapping from model
func (model *Book) MapUpdateFrom(from *Book) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"title":       model.Title,
			"isbn":        model.ISBN,
			"author":      model.Author,
			"genre":       model.Genre,
			"publisher":   model.Publisher,
			"edition":     model.Edition,
			"description": model.Description,
			"updated_at":  model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.ISBN != from.ISBN {
		mapped["isbn"] = from.ISBN
	}

	if model.Title != from.Title {
		mapped["title"] = from.Title
	}

	if model.Author != from.Author {
		mapped["author"] = from.Author
	}

	if model.Genre != from.Genre {
		mapped["genre"] = from.Genre
	}

	if model.Publisher != from.Publisher {
		mapped["publisher"] = from.Publisher
	}

	if model.Edition != from.Edition {
		mapped["edition"] = from.Edition
	}

	if model.Description != from.Description {
		mapped["description"] = from.Description
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
