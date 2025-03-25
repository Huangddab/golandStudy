package main

import "fmt"

func main() {
	var num int = 1
	var i = 11
	var prt *int = &i // 指针类型 值就是i的地址
	prt = &num        // 指针类型 值就是num的地址
	fmt.Println(prt)  // 0xc00008c098

	// 写一个程序,获取一个变量num1的地址,并显示在终端
	var num1 int = 333
	var ptr *int = &num1
	fmt.Println(ptr) // 0xc00000a108

	//练习1:
	//假如还有99天放假,问:xx星期零xx天
	date := 99
	weeks := 99 / 7
	days := date % 7

	fmt.Printf("%d 星期零 %d 天", weeks, days)

	//练习2:
	//定义一个变量保存华氏温度,华氏温度转换为摄氏温度的公式为:5 / 9 * (华氏温度 - 100),请求华氏温度对应的摄氏温度
	fmt.Println("")
	var f float32 = 123.2
	c := 5.0 / 9.0 * (f - 100)
	fmt.Printf("%.2f", c)

}
