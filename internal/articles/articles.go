package articles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	articleService := &ArticleService{db}
	articleService.db.AutoMigrate(&Article{})

	r.Get("/", articleService.GetArticles)
	r.Get("/{id}", articleService.GetSingleArticle)
	r.Post("/", articleService.CreateArticle)
	r.Put("/{id}", articleService.UpdateArticle)
	r.Delete("/{id}", articleService.DeleteArticle)

	return r
}

type Article struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ArticleService struct {
	db *gorm.DB
}

func (a *ArticleService) GetArticles(w http.ResponseWriter, r *http.Request) {
	var articles []Article
	a.db.Find(&articles)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func (a *ArticleService) GetSingleArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var article Article
	a.db.First(&article, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&article)
}

func (a *ArticleService) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "couldn't parse article")
	}
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "`id` should be a number")
	}

	article.ID = uint(id)

	a.db.Save(&article)
}

func (a *ArticleService) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var newArticle Article
	json.NewDecoder(r.Body).Decode(&newArticle)

	a.db.Create(&newArticle)

	w.WriteHeader(http.StatusCreated)
}

func (a *ArticleService) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	a.db.Delete(&Article{}, id)
}
