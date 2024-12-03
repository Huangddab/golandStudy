package main

import "fmt"

/*
panic/recover

Go 语言中目前（Go1.12）是没有异常机制，但是使用 panic/recover 模式来处理错误。

panic 可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效
*/

func fn1(a int, b int) {
	defer func() {
		// recover()必须搭配defer使用,捕获可能发生的 panic
		err := recover()
		// 如果程序出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("error:", err) //error: runtime error: integer divide by zero
		}
	}()
	fmt.Println(a / b)
}
func main() {
	fn1(10, 0)
	fmt.Println("结束")
	fn1(10, 2)
}
