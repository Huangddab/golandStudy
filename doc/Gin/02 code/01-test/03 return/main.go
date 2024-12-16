package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 1 string
	r.GET("/string", func(c *gin.Context) {
		aid := c.Query("aid")
		c.String(200, "aid=%s", aid)
	})
	// curl -X GET "http://localhost:8080/string?aid=1234"
	// aid=1234

	// 2 json
	// 方法一
	r.GET("/json", func(c *gin.Context) {
		aid := c.Query("aid")
		c.JSON(http.StatusOK, gin.H{
			"aid": aid,
		})
	})
	// curl -X GET "http://localhost:8080/json?aid=1234"
	// {"aid":"1234"}

	// 方法二：使用结构体
	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "IT 营学院"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	// curl -X GET "http://localhost:8080/moreJSON
	// {"user":"IT 营学院","Message":"Hello world!","Age":18}

	r.GET("/jsonp", func(c *gin.Context) {
		aid := c.Query("aid")
		c.JSONP(http.StatusOK, gin.H{
			"aid": aid,
		})
	})
	// curl -X GET "http://localhost:8080/jsonp?aid=1234"
	// {"aid":"1234"}

	r.Run(":8080")
}
