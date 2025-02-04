package main

import (
	"log"
	"net/http"

	"github.com/Takahiro800/go-intermediate/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/hello", handlers.HelloHandler)
	r.Post("/article", handlers.PostingArticleHandler)
	r.Get("/article/list", handlers.ArticleListHandler)
	r.Get("/article/{articleID}", handlers.ArticleShowHandler)
	r.Post("/article/nice", handlers.PostingNiceHandler)
	r.Post("/article/comment", handlers.PostingCommentHandler)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
