package login

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginForm struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login",funrouter(ctx *gin.Context){
		// 显示绑定声明绑定	multipart form
		// c.shouldBingWith(&form, binding.Form) // 显示绑定
		// 显示绑定声明绑定	json
		// c.shouldBingWith(&form, binding.JSON) // 显示绑定
		var form LoginForm
		if  c.shouldBing(&form) == nil{
			if form.User == "user"&& form.Password == "password" {
				c.JSON(200,gin.H{"status":"you are logged in"})
			}else{
				c.JSON(500,gin.H{"status":"unauthorized"})
			}
		}

	})
	router.Run(":8000")
}
