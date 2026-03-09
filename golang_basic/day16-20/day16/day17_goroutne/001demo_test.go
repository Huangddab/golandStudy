package day17

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	// 	go func() {
	// 		for i := 0; i < 5; i++ {
	// 			fmt.Println("goroutine 1", i)
	// 		}
	// 	}()

	// 	for i := 0; i < 4; i++ {
	// 		// 让出时间片，先让别的协议执行，它执行完，再回来执行此协程
	// 		runtime.Gosched()
	// 		fmt.Println("main goroutine", i)
	// 	}

	// 	// 获取go的根目录
	// 	fmt.Println("goroot:", runtime.GOROOT())
	// 	// goroot: D:\Tools\Go

	// 	fmt.Println("获取CPU数量:", runtime.NumCPU())
	// 	// 获取CPU数量: 20

	// 	fmt.Println("")

	go func() {
		// 创新的协程
		fmt.Println("goroutine starting")
		fun()
		fmt.Println("goroutine finished")
	}()

	// 阻塞 等待协程执行完成·
	time.Sleep(5 * time.Second)
}

func fun() {
	defer fmt.Println("defer fun")
	// 终止此函数 所在的协程
	runtime.Goexit()
	fmt.Println("fun")
}
