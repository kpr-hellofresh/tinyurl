package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kpr-hellofresh/tinyurl/internal/app"
)

func NewRouter(srv app.Service) http.Handler {
	router := chi.NewRouter()

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HelloWorld")
	})

	return router
}
