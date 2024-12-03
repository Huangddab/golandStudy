package main

import (
	"fmt"
	"itying/calc"
	"itying/tools"
)

func main() {

	sum := calc.Add(10, 2)
	fmt.Println(sum) //12
	// fmt.Println(calc.aaa)
	sub := calc.Sub(10, 2)
	fmt.Println(sub)

	//tools包
	b := tools.Mul(2, 6)
	fmt.Println(b)

	tools.PrintInfo()

	tools.SortIntAsc()
}

//main包中init函数 优先于 main函数
func init() {
	fmt.Println("main init...")
}
