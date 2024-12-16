package main

import "github.com/gin-gonic/gin"

func main() {
	// 1. 初始化路由 创建一个默认的路由引擎
	router := gin.Default()
	// 2. 绑定路由规则，执行的函数
	router.GET("/", func(c *gin.Context) { // 定义请求的路径和处理函数
		c.JSON(200, gin.H{ // 定义返回的JSON数据
			"message": "hello world", // 定义返回的数据
		})
	})
	// 3. 监听端口，默认在8080
	router.Run(":8080")

	// 4. 启动服务
	// 终端执行 go run main.go

	// 5. 访问服务
	// 在浏览器访问localhost:8080

	// 输出 {"message":"hello world"}
}
