package usecase

import (
	"context"

	"github.com/kien6034/secure-genom/internal/domain/repository"
)

type Service struct {
	dataUploadRepo repository.DataUploadRepository
}

func NewService(
	dataUploadRepo repository.DataUploadRepository,
) *Service {
	return &Service{
		dataUploadRepo: dataUploadRepo,
	}
}

func (s *Service) UploadData(ctx context.Context, data []byte) error {
	// pre-process data

	// store uploaded data
	s.dataUploadRepo.StoreEncryptedData(ctx, data)

	// post-process data

	return nil
}
