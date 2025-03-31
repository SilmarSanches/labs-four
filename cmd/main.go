package main

import (
	"fmt"
	_ "labs-four/docs"
	"labs-four/internal/infra/web/webserver"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Rate-Limiter
// @version 1.0
// @description Rate-Limiter
// @BasePath /
func main() {
	hello := NewGetHelloHandler()

	httpServer := webserver.NewWebServer(NewConfig())

	httpServer.AddHandler("GET", "/swagger/*", httpSwagger.WrapHandler)
	httpServer.AddHandler("GET", "/hello", hello.HandleHello)
	fmt.Println("HTTP server running at port 8080")
	httpServer.Start()
}
