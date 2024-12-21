package main

import (
	"net/http"

	"github.com/daalfox/gophlog/internal/gophlog"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/articles", gophlog.Router())
	http.ListenAndServe(":8080", r)
}
