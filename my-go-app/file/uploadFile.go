package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

/*
上传文件经过如下 3 步：
1. 在 Web 页面选择一个文件，然后上传
2. 在服务端读取上传文件的数据（字节流
3. 将文件数据写到服务端的某一个文件中
*/

func uploadFileHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024)              // 解析表单数据，最大 1MB
	file, handler, err := r.FormFile("uploadfile") // 获取上传的文件

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()                   // 延迟关闭文件（在uploadFile函数结束时关闭文件）
	fmt.Fprintf(w, "%v", handler.Header) // 输出文件信息

	// 打开服务端文件
	f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 打开文件，只写，创建文件，权限为 0666

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer f.Close() // 延迟关闭文件（在uploadFile函数结束时关闭文件）

	// 读取上传的文件数据，并写入到服务端的文件中
	io.Copy(f, file) // 拷贝文件数据
}

func showUploadfilePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./uploadFile.tpl")
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", showUploadfilePage)
	http.HandleFunc("/upload", uploadFileHandle)

	fmt.Println("服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/")

	// 启动 HTTP 服务，并监听端口号，开始监听，处理请求，返回响应
	err := http.ListenAndServe(":8900", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
