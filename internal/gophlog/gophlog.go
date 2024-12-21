package gophlog

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Router() http.Handler {
	r := chi.NewRouter()

	db, err := gorm.Open(sqlite.Open("gophlog.db"))
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&Article{})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var articles []Article
		db.Find(&articles)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(articles)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var article Article
		db.First(&article, id)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&article)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var newArticle Article
		json.NewDecoder(r.Body).Decode(&newArticle)

		db.Create(&newArticle)

		w.WriteHeader(http.StatusCreated)
	})

	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		var article Article
		if err = json.NewDecoder(r.Body).Decode(&article); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "couldn't parse article")
		}
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "`id` should be a number")
		}

		article.ID = uint(id)

		db.Save(&article)
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		db.Delete(&Article{}, id)
	})

	return r
}

type Article struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}
