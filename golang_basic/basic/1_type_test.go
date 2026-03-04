package basic

import (
	"fmt"
	"testing"
)

func TestTypr(t *testing.T) {
	//  使用type来声明一个结构体
	type Person struct {
		age  int
		name string
	}
	var person Person
	fmt.Println(person)

	// 使用type来声明一个接口
	type Animal interface {
		genre() string
		color() string
	}
}
