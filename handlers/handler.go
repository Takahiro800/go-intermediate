package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	articleID, err := strconv.Atoi(chi.URLParam(req, "articleID"))
	if err != nil {
		http.Error(w, "Invalid articleID", http.StatusBadRequest)
		return
	}

	res := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, res)
}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostingCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
