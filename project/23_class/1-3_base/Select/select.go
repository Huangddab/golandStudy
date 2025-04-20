package main

import (
	"fmt"
	"strconv"
	"time"
)

func sample1(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "hello goroutines! num:" + strconv.Itoa(i) // 向channel中写入数据 strconv.Itoa(i) 将int类型转换为string类型
		time.Sleep(1 * time.Second)
	}
	// close(ch) // 关闭channel
}

// sample1 函数向字符串通道 ch1 发送数据，数据格式为"hello goroutines! num:"加上数字。它每秒发送一次，共发送5次。

func sample2(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
	close(ch) // 关闭channel
}

// sample2 函数向整数通道 ch2 发送数据，每次发送一个整数，共发送15次。它每两秒发送一次，并在发送完毕后关闭通道。

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)
	for i := 0; i < 10; i++ {
		go sample1(ch1)
		go sample2(ch2)
	}

	for i := 0; i < 1000; i++ {
		select {
		case str, ch1check := <-ch1: // 从channel中读取数据
			/**
			* ch1check 表示channel是否关闭
			* 如果channel关闭，ch1check的值为false
			* 如果channel没有关闭，ch1check的值为true
			**/
			if !ch1check {
				fmt.Println("ch1check is false")
			}
			fmt.Println(str) // 打印channel中的数据 hello goroutines!
		case num, ch2check := <-ch2:
			if !ch2check {
				fmt.Println("ch2check is false")
			}
			fmt.Println(num)
		}
	}

	time.Sleep(60 * time.Second)
}
