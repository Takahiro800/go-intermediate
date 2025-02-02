package main

import (
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

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postingArticleHandler)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
