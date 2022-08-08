package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func showIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)

	router.Run()
}
