package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "post",
		})
	})

	r.PUT("/put", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "put",
		})
	})

	r.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "delete",
		})
	})

	// GET传值
	r.GET("/get", func(ctx *gin.Context) {
		// 获取get参数
		name := ctx.Query("name")
		age := ctx.Query("age")
		ctx.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// curl -X GET "http://localhost:8080/get?name=张三&age=18"

	// POST传值
	r.POST("/post/value", func(ctx *gin.Context) {
		// 获取post参数
		name := ctx.PostForm("name")
		age := ctx.PostForm("age")
		ctx.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
		fmt.Println("name:", name, "age", age)
	})
	// curl -d "name=张三&age=12" -X POST http://localhost:8080/post/value

	// 动态路由
	r.GET("/user/:name/:age", func(ctx *gin.Context) {
		// 获取动态路由参数
		name := ctx.Param("name")
		age := ctx.Param("age")
		ctx.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
		fmt.Println("name:", name, "age", age)
	})
	// curl -X GET "http://localhost:8080/user/张三/18"

	r.Run(":8080")
}
