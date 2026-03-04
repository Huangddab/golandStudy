package basic

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// 切片
	var s []string
	fmt.Printf("slice::%s", s)

	s1 := make([]string, 2, 2)
	fmt.Printf("slice的长度:%d", len(s1))

}
 