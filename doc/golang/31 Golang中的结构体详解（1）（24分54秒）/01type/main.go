package main

import "fmt"

//自定义类型
type myInt int

// type myFn func(int, int) int

//类型别名
type myFloat = float64

func main() {

	var a myInt = 10

	fmt.Printf("%v %T\n", a, a) //10 main.myInt

	var b myFloat = 12.3
	fmt.Printf("%v %T", b, b) //12.3 float64
}
