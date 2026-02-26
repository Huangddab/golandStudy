package main

import "fmt"

// 关键字iota 累加1 默认值是0
const (
	BEIJING  = iota + 1
	SHANGHAI = iota * 2
	JIEYANG
	SHANTOU
)

func main() {
	// 常量 只读属性

	const length int = 10

	fmt.Println("汕头=", SHANTOU)
}
