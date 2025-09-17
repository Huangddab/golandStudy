package main

import (
	"fmt"
	"sync"
	"time"
)

// 1.互斥锁

// 保证统一时间只有一个goroutine能访问共享资源

var counter int
var mutex sync.Mutex //使用 sync.Mutex，通过 Lock()和 Unlock()方法来保护临界区。
var wg sync.WaitGroup

func increment() {
	defer wg.Done() // 减少计数 -1
	mutex.Lock()    // 获取互斥锁 所有锁已经被其他gorountine持有 那么当前goroutine就会被阻塞 知道锁被释放
	counter++
	mutex.Unlock() // 释放互斥锁 允许其他被阻塞的gorountine 获取锁并访问counter
}

// 2.读写锁 允许多个
// goroutine 同时读取共享资源，但写操作是独占的（与读操作和其他写操作互斥）。这在​​读多写少​​的场景能显著提高性能。

var (
	data    map[string]string = make(map[string]string)
	rwMutex sync.RWMutex
)

// 模拟读取数据操作
func readValue(key string) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	fmt.Printf("Read:key=%s,value=%s \n", key, data[key])
	time.Sleep(10 * time.Millisecond)
}

// 模拟写入数据操作
func writeValue(key, value string) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	data[key] = value
	fmt.Printf("Write: key=%s, value=%s\n", key, value)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1) // 计数 +1
		go increment()
	}
	wg.Wait()
	fmt.Println("counter:", counter)

	// --------------------------------------

	var wgg sync.WaitGroup
	writeValue("name", "Alice")
	writeValue("name", "NEW hello")

	fmt.Println("--启动并发操作--")
	// 启动多个读goroutine
	for i := 0; i < 5; i++ {
		wgg.Add(1)
		go func(id int) {
			defer wgg.Done()
			fmt.Printf("==启动并发读操作== Goroutine %d trying to read 'name'\n", id)
			readValue("name")
			fmt.Printf("==启动并发读操作结束== Goroutine %d finished reading 'name'\n", id)
		}(i)
	}
	// 启动一个写goroutine
	wgg.Add(1)
	go func() {
		defer wgg.Done()
		fmt.Println("==启动一个写操作==Goroutine trying to write 'name' to 'Bob'")
		writeValue("name", "Bob") // 尝试修改 "name" 的值
		fmt.Println("==启动一个写操作结束==Goroutine finished writing 'name'")
	}()

	// 启动另一个读 goroutine，可能读取到新值
	wgg.Add(1)
	go func() {
		defer wgg.Done()
		time.Sleep(200 * time.Millisecond) // 等待一段时间，让写操作有机会完成
		fmt.Println("==启动一个读操作==Goroutine trying to read 'name' again")
		readValue("name")
		fmt.Println("==启动一个读操作结束==Goroutine finished reading 'name' again")
	}()

	wgg.Wait() // 等待所有 goroutine 完成
	fmt.Println("\n--- 所有操作完成 ---")
	fmt.Printf("Final data: %v\n", data)
}
