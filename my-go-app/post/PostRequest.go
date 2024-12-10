package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// 处理表单：获取HTTP Post请求字段
// 表单提交用户登录信息

func PostInfoHandler(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd() // 获取当前工作目录，并处理两个返回值
	if err != nil {
		fmt.Println("Error:", err) // 处理错误
		return
	}
	fmt.Println("Current working directory:", dir) // 输出当前工作目录
	// get请求就解析url传递的参数，POST则解析响应包的主体
	// 解析URL参数
	r.ParseForm()
	fmt.Println("r.Form", r.Form)    // 这些信息是输出到服务器端的打印信息
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		// 显示静态登录界面
		t, _ := template.ParseFiles("./login.tpl")
		t.Execute(w, nil) // 解析模板
	} else {
		// 处理post登录逻辑
		// 从r.Form中获取键为name的值

		username := ""
		password := ""
		if len(r.Form["username"][0]) > 0 {
			username = r.Form["username"][0]
		}

		if len(r.Form["password"][0]) > 0 {
			password = r.Form["password"][0]
		}

		if username == "admin" && password == "123456" {
			fmt.Fprintf(w, "登录成功！")
		} else {
			fmt.Fprintf(w, "用户名或密码错误！")

		}
	}
}
func main() {

	// 注册路由和路由函数，将url规则与处理器函数绑定做一个map映射存起来，
	// 并且会实现ServeHTTP方法，使处理器函数变成Handler函数
	http.HandleFunc("/", PostInfoHandler) // 设置访问的路由
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/login")

	// 监听并设置端口，默认是nil
	err := http.ListenAndServe(":8900", nil) // 设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
