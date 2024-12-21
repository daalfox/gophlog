package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := gorm.Open(sqlite.Open("gophlog.db"))
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	addRoutes(r, db)

	http.ListenAndServe(":8080", r)
}
