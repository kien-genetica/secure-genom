package wire

import (
	"github.com/google/wire"
	"github.com/kien6034/secure-genom/internal/domain/repository"
	"github.com/kien6034/secure-genom/internal/usecase/organization"
	"github.com/kien6034/secure-genom/internal/usecase/upload"
)

type App struct {
	UploadService *upload.Service
	OrgService    *organization.Service
}

var ProviderSet = wire.NewSet(
	upload.NewService,
	organization.NewService,
)

func InitializeApp(
	dataUploadRepo repository.DataUploadRepository,
	orgRepo repository.OrganizationRepository,
) (*App, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
