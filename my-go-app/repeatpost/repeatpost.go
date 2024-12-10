package main

import (
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
	"time"
)

/*
防止多次重复提交表单（POST 请求）

// 在表单代码放置一个隐藏字段 页面每次加载后，服务端都会为其生成一个唯一的 Token。如果发现 Token 重复提交，那么就不会处理提交上来的数据，会直接报错。
// 生成 Token 的方法：
// 1. 生成一个随机数，作为 Token 的种子。
// 2. 将随机数和当前时间戳一起写入到 hash 中。
// 3. 将 hash 转换为字符串，作为 Token。

// 在表单中放置一个隐藏字段，用于存储 Token。
// 在表单提交时，将 Token 作为参数传递给服务端。
*/

var count = 0

func repeatpostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 分析客户端的body数据
	if r.Method == "GET" {
		// 生成 Token
		currentTime := time.Now().Unix() // 时间戳
		fmt.Println("当前时间戳：", currentTime)
		h := sha512.New()                                     // 生成一个hash对象
		io.WriteString(h, strconv.FormatInt(currentTime, 10)) // 将时间戳写入hash对象
		token := fmt.Sprintf("%x", h.Sum(nil))                // 生成 Token 是一个 128 位的字符串
		fmt.Println("token：", token)

		// 通过模版装载页面
		t, _ := template.ParseFiles("./repeatpost.tpl")
		t.Execute(w, token)
	} else {
		token := r.Form.Get("token") // 从请求中获取 Token
		if token != "" {             // 校验 Token
			fmt.Println("校验 Token，持有该 Token 的页面只能提交一次")
			if count > 0 { // 如果 count 大于 0，说明已经提交过一次了
				fmt.Fprintln(w, "您重复提交了")
				return
			}
			fmt.Println(count)
			count++
		} else {
			fmt.Fprintf(w, "Token 不存在")
		}

		code := r.Form.Get("code")
		reg, _ := regexp.Compile(`<script[^>]*> </script>`)
		text := reg.ReplaceAllLiteralString(code, "")
		t, _ := template.New("test").Parse(`<html>{{ . }}</html>`)
		t.ExecuteTemplate(w, "test", text)
	}
}

func main() {
	// 注册路由和路由函数，将url规则与处理器函数绑定做一个map映射存起来，并且会实现ServeHTTP方法，使处理器函数变成Handler函数
	http.HandleFunc("/repeatpost", repeatpostHandler)

	fmt.Println("服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/repeatpost")

	// 启动 HTTP 服务，并监听端口号，开始监听，处理请求，返回响应
	err := http.ListenAndServe(":8900", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
