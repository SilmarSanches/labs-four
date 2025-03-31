//go:build wireinject
// +build wireinject

package main

import (
	"labs-four/config"
	"labs-four/internal/infra/web"
	"labs-four/internal/infra/web/webserver"
	"labs-four/internal/infra/ratelimit"
	"labs-four/internal/usecases"

	"github.com/google/wire"
)

var ProviderConfig = wire.NewSet(config.ProvideConfig)

var ProviderUseCase = wire.NewSet(
	usecases.NewHelloUseCase,
	wire.Bind(new(usecases.HelloUseCaseInterface), new(*usecases.HelloUseCase)),
)

var ProviderHandler = wire.NewSet(web.NewGetHelloHandler)

var ProviderRateLimiter = wire.NewSet(
	ratelimit.NewRedisLimiter,
	wire.Bind(new(ratelimit.RateLimitInterface), new(*ratelimit.RedisLimiter)),
)

func NewConfig() *config.AppSettings {
	wire.Build(ProviderConfig)
	return &config.AppSettings{}
}

func NewGetHelloHandler() *web.GetHelloHandler {
	wire.Build(ProviderConfig, ProviderUseCase, ProviderHandler)
	return &web.GetHelloHandler{}
}

func NewRateLimiter(conf *config.AppSettings) ratelimit.RateLimitInterface {
	rl, _ := ratelimit.NewRedisLimiter(*conf)
	return rl
}

func NewWebServer(conf *config.AppSettings, rl ratelimit.RateLimitInterface) *webserver.WebServer {
	wire.Build(webserver.NewWebServer)
	return &webserver.WebServer{}
}
