package main

import (
	"fmt"
)

// new 和 make

func main() {
	// 错误的写法
	// var userinfo map[string]string
	// userinfo["username"] = "张三"
	// fmt.Println(userinfo)
	// map 是引用类型，零值为 nil ，需要初始化才能使用

	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	fmt.Println(userinfo)

	// 切片也是引用类型
	var a = make([]int, 4, 4)
	a[0] = 1
	fmt.Println(a)

	// 错误的写法
	// var b *int
	// *b = 100
	// fmt.Println(*b)
	// 指针也是引用类型

	// new 方法给指针变量分配内存
	var b *int
	b = new(int)
	*b = 100
	fmt.Printf("%T\t %v\t %v\n", b, b, *b)
	// *int     0xc00010c0f0    100

	var f = new(bool)
	fmt.Println(*f)
	// false

}
