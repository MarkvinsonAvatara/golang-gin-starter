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
type BookMasterCreator struct {
	cfg            config.Config
	createBookRepository repository.CreateBookRepositoryUseCase
	finderBookRepository repository.FinderBookRepositoryUseCase
	cloudStorage   interfaces.CloudStorageUseCase
}

// MasterCreatorUseCase is a use case for the Master creator
type BookMasterCreatorUseCase interface {
	CreateBook(ctx context.Context, isbn int64, title string, author string, genre string, publisher string, edition int64, description string) (*entity.Book, error)
}

// NewMasterCreator creates a new MasterCreator
func BookNewMasterCreator(
	cfg config.Config,
	createBookRepository repository.CreateBookRepositoryUseCase,
	finderBookRepository repository.FinderBookRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *BookMasterCreator {
	return &BookMasterCreator{
		cfg:            cfg,
		createBookRepository: createBookRepository,
		finderBookRepository: finderBookRepository,
		cloudStorage:   cloudStorage,
	}
}

func (masterCreator *BookMasterCreator) CreateBook(ctx context.Context, isbn int64, title string, author string, genre string, publisher string, edition int64, description string) (*entity.Book, error) {
	title = strings.TrimSpace(title)
	trimmed := strings.TrimSpace(strings.ReplaceAll(title, " ", ""))

	if trimmed == "" {
		return nil, errors.ErrEmptyData.Error()
	}

	count, err := masterCreator.finderBookRepository.FindBookByTitle(ctx, trimmed)

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
	if err := masterCreator.createBookRepository.CreateBook(ctx, book); err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return book, nil

}
