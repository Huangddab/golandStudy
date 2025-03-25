package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "首页")
	})
	r.Run(":8080")
}
