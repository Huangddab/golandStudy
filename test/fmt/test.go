package main

import "fmt"

func test() {
	fmt.Printf("22")
	// 一次声明多个变量
	var (
		// 变量名  类型
		name string
	)
	name = "张三"
	fmt.Print(name)

	// const (
	// 	a = iota
	// 	b = iota
	// )
	// fmt.Print(a, b)
}
