package main

import (
	"net/http"

	"github.com/kpr-hellofresh/tinyurl/internal/platform/app"
)

func main() {
	router := app.NewRouter()

	http.ListenAndServe(":8080", router)
}
