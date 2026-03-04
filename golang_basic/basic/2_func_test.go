package basic

import (
	"fmt"
	"testing"
)

func TestFuc(t *testing.T) {
	funcTest()
}
func funcTest() {
	expression := 10
	condition := 1
	// 函数
	switch expression {
	case condition:
		fmt.Println("1111")
	default:
		fmt.Println("else")
	}

}
