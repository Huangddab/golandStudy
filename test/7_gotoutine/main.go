package main

import (
	"fmt"
	"time"
)

// 并行执行需求：
// 在主线程(可以理解成进程)中，开启一个 goroutine,
// 2.该协程每隔 50 毫秒秒输出 "你好 golang"
//
//	1.在主线程中也每隔 50 毫秒输出"你好 golang", 输出 10 次后，退出程序，
//
// 要求主线程和goroutine 同时执行。

// var wg sync.WaitGroup

// 2.协程
// func test() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("test", i+1, "你好 golang")
// 		time.Sleep(time.Millisecond * 100)
// 		// 当主线程的事件比协程时间短,主线程执行完毕协程还没执行完毕，协程就会被销毁
// 	}
// 	wg.Done() // 当协程执行完毕，计数器减1
// 	fmt.Println("test 协程退出")
// }
// func test2() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("test2", i+1, "你好 golang")
// 		time.Sleep(time.Millisecond * 100)
// 		// 当主线程的事件比协程时间短,主线程执行完毕协程还没执行完毕，协程就会被销毁
// 	}
// 	wg.Done() // 当协程执行完毕，计数器减1
// 	fmt.Println("test2 协程退出")
// }

// // 1. 主线程
// func main() {
// 	// test() // 按顺序(并行)执行test 1-10 main 1-10
// 	wg.Add(1) // 开启协程前，先让wg 协程计数器加1
// 	go test() // 开启协程(串行)
// 	wg.Add(1)
// 	go test2()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("main", i+1, "你好 golang")
// 		time.Sleep(time.Millisecond * 50)
// 	}
// 	wg.Wait() // 主线程更快时，会等待协程执行完毕，再退出
// 	fmt.Println("主线程退出")
// }
// 设置 Golang 并行运行的时候占用的 cup 数量

// func main() {
// 	// 获取当前计算机上面的CPU个数
// 	num := runtime.NumCPU()
// 	fmt.Println("cpu num=", num)
// 	// 可以自己设置使用多个cpu
// 	runtime.GOMAXPROCS(num - 1)
// 	fmt.Println("ok")
// }

// 启动多个 Goroutine
// var wg sync.WaitGroup

// func hello(i int) {
// 	defer wg.Done() // goroutine 结束就登记-1
// 	fmt.Println("Hello Goroutine!", i)
// }
// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1) // 启动一个 goroutine 就登记+1
// 		go hello(i)
// 	}
// 	wg.Wait() // 等待所有登记的 goroutine 都结束
// }

// 需求：要统计 1-120000 的数字中那些是素数？

func main() {
	// 方式一
	// 开始时间戳
	start := time.Now().Unix()
	for i := 2; i < 120000; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(i, "是素数")
		}
	}
	// 结束时间戳
	end := time.Now().Unix()
	fmt.Printf("执行耗时=%d秒", end-start)
	// 执行耗时=5秒
}
