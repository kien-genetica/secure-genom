package main

import (
	"context"
	"log"

	"github.com/kien6034/secure-genom/cmd/wire"
	"github.com/kien6034/secure-genom/internal/domain/entity"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage/adapter"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage/memory"
)

func main() {
	// Initialize ID getters
	dataIDGetter := adapter.NewDataIDGetter()

	// Initialize memory storage
	dataStorage := memory.NewMemoryStorage[*entity.Data](dataIDGetter)

	// Wire the application with the repository
	app, err := wire.InitializeApp(dataStorage)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	app.UploadService.UploadData(context.Background(), []byte("test"))
}
