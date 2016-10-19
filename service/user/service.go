package user

type Service struct {
}

func New() *Service {
	return newService()
}

func newService() *Service {
	service := &Service{}
	return service
}
