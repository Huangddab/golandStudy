package basic

import (
	"fmt"
	"testing"
	"time"
)

// select 多路复用、超时控制 与goroutine和channel组合使用

func TestSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "message2"
	}()

	for {
		select {
		case message1 := <-ch1:
			fmt.Println("received", message1)
		case message2 := <-ch2:
			fmt.Println("received", message2)

		}
	}
}
