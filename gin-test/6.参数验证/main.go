package main

import "github.com/gin-gonic/gin"


// binding:"required"不能为空不能不传
type GignUserInfo struct {
	Name       string `json:"name" binding:"required"`
	Age        int    `json:"age"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		var userInfo GignUserInfo
		err := c.ShouldBindJSON(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "success!",
			"res": userInfo,
		})
	})
	router.Run(":8000")
}
