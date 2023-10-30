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
type BookMasterUpdater struct {
	cfg            config.Config
	updaterBookRepository repository.UpdaterBookRepositoryUseCase
	cloudStorage   interfaces.CloudStorageUseCase
}

// MasterUpdaterUseCase is a use case for the Master creator
type BookMasterUpdaterUseCase interface {
	// UpdateBook updates a book
	UpdateBook(ctx context.Context, book *entity.Book) error
}

// NewMasterUpdater creates a new MasterUpdater
func NewMasterUpdater(
	cfg config.Config,
	updaterBookRepository repository.UpdaterBookRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *BookMasterUpdater {
	return &BookMasterUpdater{
		cfg:            cfg,
		updaterBookRepository: updaterBookRepository,
		cloudStorage:   cloudStorage,
	}
}

func (masterUpdater *BookMasterUpdater) UpdateBook(ctx context.Context, book *entity.Book) error {
	if err := masterUpdater.updaterBookRepository.UpdateBook(ctx, book); err != nil {
		return errors.ErrInternalServerError.Error()
	}
	return nil
}
