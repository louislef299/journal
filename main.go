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

	router.Run()
}
