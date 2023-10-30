package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/modules/master/v1/repository"
	"github.com/google/uuid"
)

// MasterDeleter is a struct that contains all the dependencies for the Master creator
type BookMasterDeleter struct {
	cfg          config.Config
	deleterBookRepository repository.DeleterBookRepositoryUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// MasterDeleterUseCase is a use case for the Master creator
type BookMasterDeleterUseCase interface {
	DeleteBook(ctx context.Context, id uuid.UUID) error
}

// NewMasterDeleter creates a new MasterDeleter
func NewMasterDeleter(
	cfg config.Config,
	deleterBookRepository repository.DeleterBookRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *BookMasterDeleter {
	return &BookMasterDeleter{
		cfg:            cfg,
		deleterBookRepository: deleterBookRepository,
		cloudStorage:   cloudStorage,
	}
}

func (masterDeleter *BookMasterDeleter) DeleteBook(ctx context.Context, id uuid.UUID) error {
	if err := masterDeleter.deleterBookRepository.DeleteBookByID(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}

