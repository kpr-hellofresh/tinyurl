package main

import (
	"net/http"

	"github.com/kpr-hellofresh/tinyurl/internal/app"
	"github.com/kpr-hellofresh/tinyurl/internal/platform/api"
	"go.uber.org/zap"
)

func main() {
	config, err := ParseConfig()
	if err != nil {
		panic(err)
	}

	logger, err := config.CreateLogger()
	if err != nil {
		panic(err)
	}

	service := app.Service{}
	router := api.NewRouter(service)
	logger.Info("Starting HTTP server ", zap.Uint16("port", config.Http.Port))
	err = http.ListenAndServe(config.ListenerAddress(), router)
	if err != nil {
		panic(err)
	}
}
