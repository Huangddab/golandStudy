package main

import "fmt"

// 值传递

// func changeValue(p int) {
// 	fmt.Println("p=", p)
// 	p = 10
// 	fmt.Println("p=", &p)
// }

// func main() {
// 	var a int = 1
// 	changeValue(a)
// 	fmt.Println("a=", a)
// 	fmt.Println("a=", &a)
// }
// p= 1
// p= 0xc00000a100
// a= 1
// a= 0xc00000a0e8

// 引用传递
func changeValue(p *int) {
	fmt.Println("p=", p)
	*p = 10
	fmt.Println("p=", &p)
}

func main() {
	var a int = 1
	changeValue(&a)
	fmt.Println("a=", a)
	fmt.Println("a=", &a)
}

// p= 0xc0000940a8
// p= 0xc000096058
// a= 10
// a= 0xc0000940a8
