package basic

import (
	"fmt"
	"sync"
	"testing"
)

// 互斥锁：独占锁 同一时刻被加锁的对象，只能被一个协程操作

type Container struct {
	mu    sync.Mutex
	count map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count[name]++
}

func TestLock(t *testing.T) {
	s := Container{
		count: map[string]int{
			"a": 0,
			"b": 2,
		},
	}
	// 用于等待goroutine完成
	var wg sync.WaitGroup

	update := func(name string, val int) {
		defer wg.Done() // 确保在函数结束时调用Done
		for i := 0; i < val; i++ {
			s.inc(name)
		}
	}

	// 设置需要等待2个goroutine
	wg.Add(2)

	// 启动goroutine
	go update("a", 100)
	go update("b", 200)

	// 等待所有goroutine完成
	wg.Wait()

	fmt.Printf("a: %d, b: %d\n", s.count["a"], s.count["b"])
}
