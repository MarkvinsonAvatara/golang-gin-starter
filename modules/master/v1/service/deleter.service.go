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
type MasterDeleter struct {
	cfg          config.Config
	bookRepository repository.BookRepositoryUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// MasterDeleterUseCase is a use case for the Master creator
type MasterDeleterUseCase interface {
	DeleteBook(ctx context.Context, id uuid.UUID) error
}

// NewMasterDeleter creates a new MasterDeleter
func NewMasterDeleter(
	cfg config.Config,
	bookRepository repository.BookRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *MasterDeleter {
	return &MasterDeleter{
		cfg:            cfg,
		bookRepository: bookRepository,
		cloudStorage:   cloudStorage,
	}
}

func (masterDeleter *MasterDeleter) DeleteBook(ctx context.Context, id uuid.UUID) error {
	if err := masterDeleter.bookRepository.DeleteBookByID(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}

