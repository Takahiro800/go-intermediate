package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Commnet struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Commnet `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	comment1 := Commnet{
		CommentID: 1,
		ArticleID: 1,
		Message:   "This is a good article!",
		CreatedAt: time.Now(),
	}

	comment2 := Commnet{
		CommentID: 2,
		ArticleID: 1,
		Message:   "I don't think so.",
		CreatedAt: time.Now(),
	}

	Article := Article{
		ID:          1,
		Title:       "Hello, World",
		Contents:    "This is a sample article.",
		UserName:    "John",
		NiceNum:     10,
		CommentList: []Commnet{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	jsonData, err := json.Marshal(Article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
