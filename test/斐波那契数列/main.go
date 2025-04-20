// 斐波那契数列 前两个相加等于后一个数

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Fibonacci(i))
	}
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
