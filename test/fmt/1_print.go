package main

import "fmt"

func main() {
	fmt.Print("111", "aa")
	fmt.Println("111", "aa")
	fmt.Println("121212")
	// Println 会换行 两个之间有空格
	fmt.Printf("你好")
	fmt.Println("121212")

	var a int = 10
	b := 20 // 类型推导方式定义变量
	fmt.Println("a=", a, "b=", b)
	fmt.Printf("a=%v b=%v", a, b)

}
