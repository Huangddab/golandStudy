package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// (1)查询参数 ?user=zhang
func _query(c *gin.Context) {
	// user := c.Query("user")
	// if user == "" {
	// 	fmt.Print("用户没有传")
	// }
	_string, ok := c.GetQuery("user")     // 判断是否有传参 获取单个参数
	fmt.Println(_string, ok)              // 有健没值 true
	_array, ok := c.GetQueryArray("user") // 获取多个参数
	// http://127.0.0.1:8000/query?user=zhang&user=san
	fmt.Println(_array, ok) // [zhang san] true
	_map, ok := c.GetQueryMap("user")
	// http://127.0.0.1:8000/query?user[id]=123&user[name]=zhanghsan
	fmt.Println(_map, ok) // map[id:123 name:zhanghsan] true

}

// (2)动态参数 /:user_id
func _param(c *gin.Context) {
	_userID := c.Param("user_id")
	_bookID := c.Param("book_id")
	fmt.Println(_userID, _bookID) // 123 456
}

// (3) 表单参数	只能POST
func _form(c *gin.Context) {
	_username := c.PostForm("username")
	_password := c.PostForm("password") // 获取单个
	fmt.Println(_username, _password)
	// huangdan 123456

	_array, ok := c.GetPostFormArray("hobby") // 获取多个
	fmt.Println(_array, ok)                   // [篮球 羽毛球] true

	_string := c.DefaultPostForm("addr", "广东省") // 获取单个 没有返回默认值
	fmt.Println(_string)

	_forms, err := c.MultipartForm() // 获取所有form参数
	fmt.Println(_forms, err)         // &{map[name:[huangdan]] map[file:[0xc00019c120]]} <nil>

}
func main() {
	router := gin.Default()
	router.PATCH("/query", _query)

	router.POST("/param/:user_id", _param)
	// http: //127.0.0.1:8000/param/123
	router.POST("/param/:user_id/:book_id", _param)
	// http://127.0.0.1:8000/param/123/456

	router.POST("/form", _form)

	router.Run(":8000")

}
