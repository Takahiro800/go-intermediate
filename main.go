package main

import (
	"log"
	"net/http"

	"github.com/Takahiro800/go-intermediate/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostingArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleShowHandler)
	http.HandleFunc("/article/nice", handlers.PostingNiceHandler)
	http.HandleFunc("/article/comment", handlers.PostingCommentHandler)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
