package service

import (
	"context"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/modules/master/v1/repository"
	// "github.com/google/uuid"
	"gin-starter/common/errors"
	"gin-starter/entity"
)

// MasterUpdater is a struct that contains all the dependencies for the Master creator
type MasterUpdater struct {
	cfg            config.Config
	bookRepository repository.BookRepositoryUseCase
	cloudStorage   interfaces.CloudStorageUseCase
}

// MasterUpdaterUseCase is a use case for the Master creator
type MasterUpdaterUseCase interface {
	// UpdateBook updates a book
	UpdateBook(ctx context.Context, book *entity.Book) error
}

// NewMasterUpdater creates a new MasterUpdater
func NewMasterUpdater(
	cfg config.Config,
	bookRepository repository.BookRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *MasterUpdater {
	return &MasterUpdater{
		cfg:            cfg,
		bookRepository: bookRepository,
		cloudStorage:   cloudStorage,
	}
}

func (masterUpdater *MasterUpdater) UpdateBook(ctx context.Context, book *entity.Book) error {
	if err := masterUpdater.bookRepository.UpdateBook(ctx, book); err != nil {
		return errors.ErrInternalServerError.Error()
	}
	return nil
}
