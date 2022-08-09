package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)
	router.GET("/journal/view/:date", getEntry)
	router.POST("journal", addEntry)
	/*
		curl -X POST -H 'Content-Type: application/json' \
		-d '{"title":"title", "content":"content"}' http://localhost:8080/journal
	*/

	router.Run()
}
