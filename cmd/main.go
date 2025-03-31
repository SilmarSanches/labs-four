package main

import (
	"fmt"
	_ "labs-four/docs"
	"labs-four/internal/infra/ratelimit"
	"labs-four/internal/infra/web/webserver"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Rate-Limiter
// @version 1.0
// @description Rate-Limiter
// @BasePath /
func main() {
	conf := NewConfig()

	var limiter ratelimit.RateLimitInterface
	if conf.RateLimitType == "redis" {
		lim, err := ratelimit.NewRedisLimiter(*conf)
		if err != nil {
			panic(err)
		}
		limiter = lim
	} else {
		limiter = ratelimit.NewMemoryLimiter(conf.DefaultIPLimit, conf.RateLimitDuration, conf.BlockDuration)
	}

	hello := NewGetHelloHandler()
	webServer := webserver.NewWebServer(conf, limiter)

	webServer.AddHandler("GET", "/swagger/*", httpSwagger.WrapHandler)
	webServer.AddHandler("GET", "/hello", hello.HandleHello)

	fmt.Println("Servidor iniciado na porta " + conf.Port)
	webServer.Start()
}
