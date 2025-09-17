package main

import "fmt"

func main() {

	a := []int{2: 1}
	fmt.Println(a)
	x := []int{4: 44, 55, 66, 1: 77, 88}
	fmt.Println(len(x), x[2])
}
