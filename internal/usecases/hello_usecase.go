package usecases

import "labs-four/config"

type HelloUseCaseInterface interface {
	Hello() string
}

type HelloUseCase struct {
	appConfig *config.AppSettings
}

func NewHelloUseCase(appConfig *config.AppSettings) *HelloUseCase {
	return &HelloUseCase{
		appConfig: appConfig,
	}
}

func (h *HelloUseCase) Hello() string {
	return "Hello, World!"
}
