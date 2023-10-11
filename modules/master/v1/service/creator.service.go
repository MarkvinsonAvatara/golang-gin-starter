package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/master/v1/repository"
	"github.com/google/uuid"
	//"gorm.io/gorm"
	"strings"
)

// MasterCreator is a struct that contains all the dependencies for the Master creator
type MasterCreator struct {
	cfg            config.Config
	bookRepository repository.BookRepositoryUseCase
	cloudStorage   interfaces.CloudStorageUseCase
}

// MasterCreatorUseCase is a use case for the Master creator
type MasterCreatorUseCase interface {
	CreateBook(ctx context.Context, isbn int64, title string, author string, genre string, publisher string, edition int64, description string) (*entity.Book, error)
}

// NewMasterCreator creates a new MasterCreator
func NewMasterCreator(
	cfg config.Config,
	bookRepository repository.BookRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *MasterCreator {
	return &MasterCreator{
		cfg:            cfg,
		bookRepository: bookRepository,
		cloudStorage:   cloudStorage,
	}
}

func (masterCreator *MasterCreator) CreateBook(ctx context.Context, isbn int64, title string, author string, genre string, publisher string, edition int64, description string) (*entity.Book, error) {
	title = strings.TrimSpace(title)
	trimmed := strings.TrimSpace(strings.ReplaceAll(title, " ", ""))

	if trimmed == "" {
		return nil, errors.ErrEmptyData.Error()
	}

	count, err := masterCreator.bookRepository.FindBookByTitle(ctx, trimmed)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	if count > 0 {
		return nil, errors.ErrDuplicateEntry.Error()
	}

	book := entity.NewBook(
		uuid.New(),
		isbn,
		title,
		author,
		genre,
		publisher,
		edition,
		description,
		"System",
	)
	if err := masterCreator.bookRepository.CreateBook(ctx, book); err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return book, nil

}
