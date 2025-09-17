package main

import (
	"fmt"
	"time"
)

func f(form string) {
	for i := 0; i < 3; i++ {
		fmt.Println(form, ":", i)
	}
}

func main() {
	// 同步的调用
	f("direct")

	// 异步 协程调用 是以并发的形式执行的
	go f("goroutune")

	// 匿名函数启动一个协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 睡眠一秒
	time.Sleep(time.Second)
	fmt.Println("done")

}
