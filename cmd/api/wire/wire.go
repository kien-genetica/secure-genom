var ProviderSet = wire.NewSet(
	storage.NewService,
	tee.processor.NewService,
	tee.delegate.NewService,
	organization.NewService,
	data.NewService,
	validator.NewService,
	wire.Bind(new(storage.StorageInterface), new(*storage.Service)),
	wire.Bind(new(tee.ProcessorInterface), new(*tee.processor.Service)),
	wire.Bind(new(tee.DelegateInterface), new(*tee.delegate.Service)),
)

func NewApp(
	orgRepo repository.OrganizationRepository,
	dataRepo repository.DataRepository,
) *App {
	wire.Build(
		ProviderSet,
		NewHTTPServer,
		wire.Struct(new(App), "*"),
	)
	return nil
} 