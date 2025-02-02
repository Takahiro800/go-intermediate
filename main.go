package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World\n")
	}

	postingArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting article...\n")
	}

	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	articleShowHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		res := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, res)
	}

	postingNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	postingCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postingArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", articleShowHandler)
	http.HandleFunc("/article/nice", postingNiceHandler)
	http.HandleFunc("/article/comment", postingCommentHandler)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
