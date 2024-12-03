package main

import (
	"errors"
	"fmt"
)

/*
panic/recover

Go 语言中目前（Go1.12）是没有异常机制，但是使用 panic/recover 模式来处理错误。

panic 可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效
*/

// 模拟了一个读取文件的方法
func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	// 使用 defer 关键字延迟执行一个匿名函数
	defer func() {
		// 调用 recover 函数，捕获可能发生的 panic
		err := recover()
		// 如果捕获到 panic，打印错误信息
		if err != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("main.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	myFn()
	fmt.Println("继续执行...")
}
