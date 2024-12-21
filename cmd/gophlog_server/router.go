package main

import (
	"github.com/daalfox/gophlog/internal/articles"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func addRoutes(r *chi.Mux, db *gorm.DB) {
	r.Mount("/articles", articles.NewServer(db))
}
