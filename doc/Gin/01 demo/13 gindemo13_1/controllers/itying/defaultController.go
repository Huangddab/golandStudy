package itying

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	//设置sessions
	session := sessions.Default(c)
	session.Set("username", "张三 111")
	session.Save() //设置session的时候必须调用

	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788418,
	})
}
func (con DefaultController) News(c *gin.Context) {
	//获取sessions
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200, "username=%v", username)
}
