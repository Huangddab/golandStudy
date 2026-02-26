package main

import "fmt"

// 方法一二三可以生命全局变量

var gb int = 1

/*
四种变量的生命方式
*/
func main() {
	// 方法一 默认值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of = %T\n", a)

	// 方法二 初始化一个值
	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("type of = %T\n", b)

	// 方法三 自动匹配
	var c = 100
	fmt.Println("c=", c)
	fmt.Printf("type of = %T\n", c)

	// 方法四 简写
	d := "abc"
	fmt.Println("d=", d)
	fmt.Printf("type of = %T\n", d)

	fmt.Println("gb=", gb)
}
