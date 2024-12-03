package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("a的值：%v a的类型：%Ta的地址：%p\n ", a, a, &a)
	// a的值：10 a的类型：int a的地址：0xc00010c098
	var b = &a // b是指针变量，存储的是a的地址
	fmt.Printf("b的值：%v b的类型：%Tb的地址：%p\n ", b, b, &b)
	// b的值：0xc00000a0e8 b的类型：*int b的地址：0xc000078058

	

}
