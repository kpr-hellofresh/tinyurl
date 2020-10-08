package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HelloWorld")
	})

	return router
}
