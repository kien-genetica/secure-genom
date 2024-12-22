package wire

import (
	"github.com/google/wire"
	"github.com/kien6034/secure-genom/internal/domain/repository"
	usecase "github.com/kien6034/secure-genom/internal/usecase/upload"
)

type App struct {
	UploadService *usecase.UploadUseCase
}

var DataUploadUcaseSet = wire.NewSet(
	repository.NewDataUploadRepository,
	usecase.NewService,
)

func InitializeApp(
	dataUploadRepo repository.DataUploadRepository,
) (*App, error) {
	wire.Build(
		DataUploadUcaseSet,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
