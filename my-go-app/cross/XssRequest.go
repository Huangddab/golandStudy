package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func XssHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./xss.tpl")
		t.Execute(w, nil)
	} else {
		code := r.Form.Get("code")                                 // 从请求中获取表单数据
		reg, _ := regexp.Compile(`<script[^>]*> </script>`)        // 定义正则表达式
		text := reg.ReplaceAllLiteralString(code, "")              // 使用正则表达式替换表单数据
		t, _ := template.New("test").Parse(`<html>{{ . }}</html>`) // 定义模板
		t.ExecuteTemplate(w, "test", text)
	}
}

func main() {
	http.HandleFunc("/xss", XssHandler)

	fmt.Println("服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/xss")

	// 启动 HTTP 服务，并监听端口号，开始监听，处理请求，返回响应
	err := http.ListenAndServe(":8900", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
