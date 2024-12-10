package main

import (
	"fmt"
	"net/http"
)

// 处理表单：获取HTTP Get请求字段

func GetInfoHandler(w http.ResponseWriter, r *http.Request) {
	// get请求就解析url传递的参数，POST则解析响应包的主体
	// 解析URL参数
	r.ParseForm()
	fmt.Println("r.Form", r.Form) // 这些信息是输出到服务器端的打印信息
	names, ok := r.Form["name"]   // 从r.Form中获取键为name的值
	// 注意：如果ok为false，则说明没有获取到值   如果ok为true，则说明获取到了值
	if ok == true {
		fmt.Println("names[0]", names[0]) // 输出name的值
		fmt.Println("r.Form[\"age\"][0]", r.Form["age"][0])
	}
	// 遍历表单
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}
	fmt.Fprintf(w, "Hello Get!")
}

func main() {
	// 注册路由和路由函数，将url规则与处理器函数绑定做一个map映射存起来，
	// 并且会实现ServeHTTP方法，使处理器函数变成Handler函数
	http.HandleFunc("/", GetInfoHandler) // 设置访问的路由

	fmt.Println("服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/a.html?name=Bill&age=12")

	// 监听并设置端口，默认是nil
	err := http.ListenAndServe(":8900", nil) // 设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

// PS E:\workspace\Go\gostudy\my-go-app> go run getInfo.go
// 服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/a.html?name=Bill&age=12
// r.Form map[age:[12] name:[Bill]]
// names[0] Bill
// r.Form["age"][0] 12
// key: name
// val: [Bill]
// key: age
// val: [12]
// r.Form map[]
