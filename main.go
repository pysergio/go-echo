package main

import (
	"go-echo/pkg/events"
	"go-echo/pkg/routers"
	setting "go-echo/pkg/settings"

	"github.com/fasthttp/router"
	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
)

func init() {
	logger.SetLevel("debug")
}

func main() {
	setting.Setup()

	events.StartApp()
	defer events.StopApp()

	router := router.New()
	router.GET("/health", routers.Health)
	server := &fasthttp.Server{
		Name:    "Go-Echo",
		Handler: router.Handler,
	}
	logger.Debug("Listening in http://localhost:8000...")
	logger.Fatal(server.ListenAndServe(":8000"))

}
