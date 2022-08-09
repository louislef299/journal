package main

import (
	"github.com/gin-gonic/gin"

	. "github.com/louislef299/journal/article"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	// curl -X GET -H "Accept: application/json" http://localhost:8080/
	router.GET("/", ShowIndexPage)
	router.GET("/journal/view/:date", GetEntry)
	router.POST("journal", AddEntry)
	/*
		curl -X POST -H 'Content-Type: application/json' \
		-d '{"title":"title", "content":"content"}' http://localhost:8080/journal
	*/

	router.Run()
}
