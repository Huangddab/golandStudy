package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/index", func(ctx *gin.Context) {
		ctx.String(200, "Hello World")
	})

	router.Run(":8080")
}
