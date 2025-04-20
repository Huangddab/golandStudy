// 实现阶乘

package main

import "fmt"

func main() {
	fmt.Println(Factorial(15))
}

func Factorial(n uint) uint {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}
