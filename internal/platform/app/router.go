package app

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HelloWorld")
	})
}
