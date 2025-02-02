package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, "Hello, World\n")
}

func PostingArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, "Posting article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, "Article List\n")
}

func ArticleShowHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	articleID := 1
	res := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, res)
}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, "Posting Nice...\n")
}

func PostingCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, "Posting Comment...\n")
}
