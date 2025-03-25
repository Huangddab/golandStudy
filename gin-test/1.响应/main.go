package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// md:https://docs.fengfengzhidao.com/#/docs/Gin%E6%A1%86%E6%9E%B6%E6%96%87%E6%A1%A3/2.%E5%93%8D%E5%BA%94

func _string(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func _json(c *gin.Context) {
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `josn:"age_num"`
		Password string `json:"-"`
		// -代表不进行json序列化 忽略转换json
	}
	// 1.
	// user := UserInfo{"张三", 34, "1234"}
	// c.JSON(200, user)

	// 2.
	// josn响应map
	// userMap := map[string]string{
	// 	"user_name": "张三",
	// 	"age_num":   "34",
	// }
	// c.JSON(200, userMap)

	// 3.
	// 直接响应json
	c.JSON(200, gin.H{
		"user_name": "张三",
		"age_num":   "34"})

}

// 重定向
func _redirect(c *gin.Context) {
	c.Redirect(302, "http://bing.com")
}

func main() {
	router := gin.Default()
	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/bing", _redirect)

	// 带参数GET请求
	router.Run(":8000")
	// 直接访问http://127.0.0.1/

}
