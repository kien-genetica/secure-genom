package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kien6034/secure-genom/internal/domain/entity"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage/adapter"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage/memory"
)

type DataUpload struct {
	storage.Storage[*entity.Data]
}

var _ DataUploadRepository = (*DataUpload)(nil)

// NewDataUploadRepository creates a new DataUpload repository
func NewDataUploadRepository() *DataUpload {
	return &DataUpload{Storage: memory.NewMemoryStorage[*entity.Data](adapter.NewDataIDGetter())}
}

func (r *DataUpload) StoreEncryptedData(ctx context.Context, data []byte) error {
	id := uuid.New().String()
	r.Storage.Create(ctx, &entity.Data{
		ID:            id,
		EncryptedData: data,
		CreatedAt:     time.Now(),
	})

	return nil
}

func (r *DataUpload) GetEncryptedData(ctx context.Context, id string) ([]byte, error) {
	return nil, nil
}
