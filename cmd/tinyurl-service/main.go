package main

import (
	"net/http"
	"time"

	"github.com/kpr-hellofresh/tinyurl/internal/app"
	"github.com/kpr-hellofresh/tinyurl/internal/domain/url"
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

	service := app.Service{
		ShortenURL: app.ShortenURL{
			Adder:     nil, // repo
			Shortener: url.Shortener{},
			Now:       time.Now,
		},
		GetURL: app.GetURL{
			Getter: nil, // repo
		},
	}
	router := api.NewRouter(service)
	logger.Info("Starting HTTP server ", zap.Uint16("port", config.Http.Port))
	err = http.ListenAndServe(config.ListenerAddress(), router)
	if err != nil {
		panic(err)
	}
}
