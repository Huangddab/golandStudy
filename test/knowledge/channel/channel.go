package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---------------1 无缓冲channel-----------------")

	// ===1 无缓冲channel （只有对应的接收通道准备好了才能发送数据）===
	// 创建一个string类型的channel
	messages := make(chan string)

	// 在新的goroutine中向channel发送数据
	go func() {
		messages <- "ping"
	}()

	// 从channel接收数据并存储在变量中
	msg := <-messages

	// msg是从channel中取出的实际数据
	fmt.Println(msg) // 输出: ping

	// 直接打印channel会显示channel的内存地址
	// 因为channel本身是一个引用类型
	fmt.Println(messages) // 输出类似: 0xc000094060

	fmt.Println("---------------2 有缓冲channel-----------------")

	// ===2 有缓冲channel （可以发送多个数据，但超过缓冲区大小后，会阻塞）===

	messages1 := make(chan string, 2)

	messages1 <- "buffered"
	messages1 <- "channel"
	// messages1 <- "channel3" 超过 报错

	msg1 := <-messages1
	fmt.Println(msg1)

	msg1 = <-messages1
	fmt.Println(msg1)

	fmt.Println("---------------3 channel 同步-----------------")

	// ===3 channel 同步===

	done := make(chan bool, 1)
	go worker(done)

	<-done // 阻塞 直到收到done通道的通知
	// 当把 <-done 注释掉 程序会直接退出

	fmt.Println("---------------4 channel 方向-----------------")

	// ===4 channel 方向===

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	fmt.Println("---------------5 channel 选择器-----------------")

	// ===5 通道选择器===

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "result 2"
	}()

	// 同时等待这两个值被发送，哪个先被发送就先接收哪个
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("---------------6 channel 超时-----------------")

	// ===6 channel 超时===

	c3 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c3 <- "result 3"
	}()

	select {
	case res := <-c3:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c4 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c4 <- "result 4"
	}()
	select {
	case res := <-c4:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	fmt.Println("---------------7 channel 非阻塞通道操作-----------------")

	// ===7 channel 非阻塞通道操作===

	message11 := make(chan string)
	signals := make(chan bool)
	// 如果message11 有值 则接收 否则执行default
	select {
	case msg := <-message11:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// msg11不能被发送到message11 因为这是一个无缓冲通道 并且没有接受者
	msg11 := "hi"
	select {
	case message11 <- msg11:
		fmt.Println("sent message", msg11)
	default:
		fmt.Println("no message sent")
	}

	// 多路的非阻塞选择器
	select {
	case msg := <-message11:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	fmt.Println("---------------8 channel 关闭-----------------")

	// ===8 channel 关闭===

	jobs := make(chan int, 5)
	done1 := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done1 <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done1

}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	done <- true
	// done 通道被用来 通知其他协程 工作已经完成
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
