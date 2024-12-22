// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kien6034/secure-genom/internal/domain/entity"
	"github.com/kien6034/secure-genom/internal/domain/repository"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage"
	"github.com/kien6034/secure-genom/internal/usecase/upload"
)

// Injectors from wire.go:

func InitializeApp(dataStorage storage.Storage[*entity.Data]) (*App, error) {
	dataUploadRepository := repository.NewDataUploadRepository(dataStorage)
	service := usecase.NewService(dataUploadRepository)
	app := &App{
		UploadService: service,
	}
	return app, nil
}

// wire.go:

type App struct {
	UploadService *usecase.Service
}

var UploadServiceSet = wire.NewSet(repository.NewDataUploadRepository, usecase.NewService)
