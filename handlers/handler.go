package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Takahiro800/go-intermediate/models"
	"github.com/go-chi/chi/v5"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World\n")
}

func PostingArticleHandler(w http.ResponseWriter, req *http.Request) {
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "fail to read Content-Length", http.StatusBadRequest)
		return
	}
	reqBodyBuffer := make([]byte, length)

	if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to read request body", http.StatusBadRequest)
	}

	defer req.Body.Close()

	var reqArticle models.Article
	if err := json.Unmarshal(reqBodyBuffer, &reqArticle); err != nil {
		http.Error(w, "fail to decode json", http.StatusBadRequest)
		return
	}

	// article := models.Article1
	// jsonData, err := json.Marshal(article)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }
	//
	// w.Write(jsonData)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articles := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articles)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func ArticleShowHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(chi.URLParam(req, "articleID"))
	if err != nil {
		http.Error(w, "Invalid articleID", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusBadRequest)
	}

	w.Write(jsonData)
}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostingCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
