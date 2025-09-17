package main

import (
	"fmt"
	"time"
)

// https://cuihairu.github.io/hello-golang/concurrency/Actor.html
// Actor 模型 每个都是独立的 并发的实体 只有通过消息传递来进行通信 避免传统并发编程的共享状态和锁的问题

// 消息接口定义
type Message interface{}

// Actor 接口定义
type Actor interface {
	Receive(msg Message)
}

// 简单的Actor实现
type SimpleActor struct {
	id      int
	mailbox chan Message
}

// NewSimoleActor：创建Actor实例，初始化邮箱通道
func NewSimoleActor(id int) *SimpleActor {
	return &SimpleActor{
		id:      id,
		mailbox: make(chan Message),
	}
}

// Receive 接收并处理消息
func (a *SimpleActor) Receive(msg Message) {
	switch msg := msg.(type) {
	case string:
		fmt.Printf("Actor %d received message: %s\n", a.id, msg)
	case int:
		fmt.Printf("Actor %d received number: %d\n", a.id, msg)
	default:
		fmt.Printf("Actor %d received unknown message\n", a.id)
	}
}

// Send 发送消息 
func (a *SimpleActor) Send(msg Message) {
	a.mailbox <- msg
}

// Start：启动goroutine监听邮箱，处理收到的消息
func (a *SimpleActor) Start() {
	go func() {
		for msg := range a.mailbox {
			a.Receive(msg)
		}
	}()
}

func main() {
	actor1 := NewSimoleActor(1)
	actor2 := NewSimoleActor(2)

	actor1.Start()
	actor2.Start()

	actor1.Send("HELLO,1")
	actor2.Send(22)

	actor1.Send("Hello,1")
	actor2.Send("Message for Actor 2")

	// 给一些时间让 goroutines 处理消息
	time.Sleep(1 * time.Second)
}
