package main

import (
	"net/http"

	"github.com/kpr-hellofresh/tinyurl/internal/platform/app"
)

func main() {
	config, err := ParseConfig()
	if err != nil {
		panic(err)
	}

	router := app.NewRouter()

	err = http.ListenAndServe(config.ListenerAddress(), router)
	if err != nil {
		panic(err)
	}
}
