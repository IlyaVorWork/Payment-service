package payment

type IProvider interface {
}

type Service struct {
	provider IProvider
}

func NewService(provider IProvider) *Service {
	return &Service{
		provider: provider,
	}
}
