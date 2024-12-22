//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kien6034/secure-genom/internal/domain/entity"
	"github.com/kien6034/secure-genom/internal/domain/repository"
	"github.com/kien6034/secure-genom/internal/infrastructure/storage"
	usecase "github.com/kien6034/secure-genom/internal/usecase/upload"
)

type App struct {
	UploadService *usecase.Service
}

var UploadServiceSet = wire.NewSet(
	repository.NewDataUploadRepository,
	usecase.NewService,
)

func InitializeApp(
	dataStorage storage.Storage[*entity.Data],
) (*App, error) {
	wire.Build(
		UploadServiceSet,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
