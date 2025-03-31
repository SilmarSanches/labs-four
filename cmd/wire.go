//go:build wireinject
// +build wireinject

package main

import (
	"labs-four/config"
	"labs-four/internal/infra/web"
	"labs-four/internal/usecases"

	"github.com/google/wire"
)

var ProviderConfig = wire.NewSet(config.ProvideConfig)

var ProviderUseCase = wire.NewSet(
	usecases.NewHelloUseCase,
	wire.Bind(new(usecases.HelloUseCaseInterface), new(*usecases.HelloUseCase)),
)

var ProviderHandler = wire.NewSet(web.NewGetHelloHandler)

func NewConfig() *config.AppSettings {
	wire.Build(ProviderConfig)
	return &config.AppSettings{}
}

func NewGetHelloHandler() *web.GetHelloHandler {
    wire.Build(ProviderConfig, ProviderUseCase, ProviderHandler)
    return &web.GetHelloHandler{}
}