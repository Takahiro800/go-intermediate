package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World\n")
}

func PostingArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

func ArticleShowHandler(w http.ResponseWriter, req *http.Request) {
	articleID := chi.URLParam(req, "articleID")
	res := fmt.Sprintf("Article No.%v\n", articleID)
	io.WriteString(w, res)
}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostingCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
