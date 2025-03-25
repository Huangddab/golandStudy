package main

import "github.com/gin-gonic/gin"

// Should Bind
// 可以绑定json，query，param，yaml，xml

// 如果校验不通过会返回错误

// ShouldBindJSON 绑定JSON数据
// ShouldBindQuery 绑定URL参数

//JSON tag为json query tag为form  uri tag为uri
type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	router := gin.Default()
	// ShouldBindJSON 绑定JSON数据
	// http://127.0.0.1:8080/
	// {
	// "name":"黄丹",
	// "age":18,
	// "sex":"feman"
	// }
	router.POST("/", func(c *gin.Context) {
		var userInfo UserInfo              // 定义一个结构体
		err := c.ShouldBindJSON(&userInfo) // 把请求中的数据反序列化到userInfo中
		if err != nil {
			c.JSON(200, gin.H{
				"msg": err.Error(),
			})
			return
		} // 反序列化失败，返回错误信息
		c.JSON(200, gin.H{
			"msg":      "success",
			"response": userInfo,
		}) // 反序列化成功，返回成功信息
	})

	// ShouldBindQuery 绑定URL参数 结构体参数要加form
	// http://127.0.0.1:8080/query?name=zhangsan&age=18&sex=man
	router.POST("/query", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindQuery(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg":      "success",
			"response": userInfo,
		})
	})

	// ShouldBindUri 绑定URL参数 结构体参数要加uri
	// http://127.0.0.1:8080/uri/fengfeng/21/男
	router.POST("uri/:name/:age/:sex", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindUri(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg":      "success",
			"response": userInfo,
		})
	})

	// ShouldBind 绑定所有数据 formdata参数
	// 	会根据请求头中的content-type去自动绑定
	// form-data的参数也用这个，tag用form
	// 默认的tag就是form
	// 绑定form-data、x-www-form-urlencode
	router.POST("/form", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBind(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{
				"msg": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"msg":      "success",
			"response": userInfo,
		})
	})

	router.Run(":8080")
}
