package main

import (
	"fmt"
	"time"
)

/**
* FIFO 先进先出 first in first out
* 3 表示有三个缓存 {[hello goroutines! 1],[hello goroutines! 2],[hello goroutines! 3]} 就可以对channel多次写入
*
**/
func sample1(message chan string) {
	message <- "hello goroutines! 1"
	message <- "hello goroutines! 2"
	message <- "hello goroutines! 3"
    // 这里不能close channel的原因是:
    // 1. sample2函数还需要从channel中读取数据
    // 2. 如果在这里close了channel，sample2再读取数据时会发生panic
    // 3. 一般由写入方关闭channel，这里sample1和sample2都会写入，应该由最后写入的sample2来关闭
}

func sample2(message chan string) {
	time.Sleep(2 * time.Second)
	str := <-message
	str += "I'm a goroutines"
	message <- str
	close(message) // 关闭channel
}

func main() {
	var message = make(chan string, 3)
	go sample1(message)
	go sample2(message)
	time.Sleep(3 * time.Second)
	// fmt.Println(<-message)
	// fmt.Println(<-message)
	// fmt.Println(<-message)
	for str := range message { // 遍历channel message为空时range不会结束 channel 关闭时range才会结束
		fmt.Println(str)
	}
	fmt.Println("hello world!")
}

// func main() {
// message := make(chan string)
// 开启一个goroutine
// go func() {
// 	message <- "hello goroutines!"
// 	// fmt.Println("hello goroutines!")
// }()
// go func() {
// 	time.Sleep(2 * time.Second)
// 	str := <-message
// 	str += "I'm a goroutines"
// 	message <- str
// }()
// 	time.Sleep(3 * time.Second)
// 	fmt.Println(<-message)
// 	fmt.Println("hello world!") // 标准输出 stdin 标准输入 stdout 标准输出
// }

// hello goroutines!I'm a goroutines
// hello world!
