package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", func(ctx *gin.Context) {

	})
	router.Run(":8000")
}
