package main

import (
	"fmt"
	"time"
)

// 定义事件类型
type Event struct {
	Type    string
	Payload interface{}
}

func main() {
	// 创建一个事件channel
	eventChan := make(chan Event)

	// 事件生产者
	go func() {
		for {
			eventChan <- Event{Type: "user_created", Payload: "user123"}
			time.Sleep(time.Second)
		}
	}()

	// 事件消费者
	go func() {
		for event := range eventChan {
			fmt.Printf("Received event: %s with payload: %v\n", event.Type, event.Payload)
		}
	}()

	// 保持程序运行
	select {}
}
