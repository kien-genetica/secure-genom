package main

import (
	"log"

	"github.com/kien6034/secure-genom/cmd/wire"
	"github.com/kien6034/secure-genom/internal/domain/entity"
	"github.com/kien6034/secure-genom/internal/domain/repository"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage/adapter"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage/memory"
)

func main() {
	// Initialize ID getters
	dataIDGetter := adapter.NewDataIDGetter()

	// Initialize memory storage
	dataStorage := memory.NewMemoryStorage[*entity.Data](dataIDGetter)

	// Initialize repositories
	dataUploadRepo := repository.NewDataUploadRepository(dataStorage)

	// Wire the application
	_, err := wire.InitializeApp(dataUploadRepo)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
}
