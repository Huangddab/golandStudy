package basic

import (
	"fmt"
	"testing"
)

// 无缓冲 同步通信
// 有缓冲 异步队列、生产者消费者

// 赋值 ch <- 1
// 取值 <-ch

func TestChannel(t *testing.T) {
	// 1)无缓冲区
	ch := make(chan int)
	// ch <- 10
	// 没有缓冲区 就会阻塞在这里 等着被消费了才会向后执行

	go func() {
		val := <-ch
		fmt.Println(val)
	}()
	ch <- 10

	// 2)有缓冲区
	ch1 := make(chan int, 10)
	ch1 <- 10

}
