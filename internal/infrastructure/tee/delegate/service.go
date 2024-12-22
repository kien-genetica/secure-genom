package delegate

import (
	"context"
)

type Service struct {
	storage storage.StorageInterface
}

func NewService(storage storage.StorageInterface) *Service {
	return &Service{storage: storage}
}

func (s *Service) ProcessReEncryption(
	ctx context.Context,
	data []byte,
	reEncryptionKey string,
) ([]byte, error) {
	// 1. Decrypt re-encryption key
	// 2. Perform proxy re-encryption
	// 3. Return re-encrypted data
	return nil, nil
}
