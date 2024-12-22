package main

import (
	"log"

	"github.com/kien6034/secure-genom/cmd/wire"
)

func main() {
	// Initialize repositories
	dataUploadRepo := s3.NewDataUploadRepository()
	orgRepo := postgres.NewOrganizationRepository()

	// Wire the application
	app, err := wire.InitializeApp(dataUploadRepo, orgRepo)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Start your application logic here
}
