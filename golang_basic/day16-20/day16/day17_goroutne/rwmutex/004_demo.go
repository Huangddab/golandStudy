package main

import (
	"fmt"
	"sync"
	"time"
)

/**
RWMutex是读/写互斥锁。锁可以由任意数量的读取器或单个编写器持有。RWMutex的零值是未锁定的mutex。
当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁； 写锁 --> 其他写锁阻塞
当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续；读锁 --> 其他读锁可用
当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定；读锁 --> 其他写锁等读锁解锁了才能写锁

所以说这里的读锁定（RLock）目的其实是告诉写锁定：有很多人正在读取数据，你给我站一边去，等它们读（读解锁）完你再来写（写锁定）

总结：
同时只能有一个 goroutine 能够获得写锁定。
同时可以有任意多个 gorouinte 获得读锁定。
同时只能存在写锁定或读锁定（读和写互斥）。

1、可以随便读，多个goroutine同时读。

2、写的时候，啥也不能干。不能读也不能写。
**/

var (
	RWMutex *sync.RWMutex   // 读写锁
	wg      *sync.WaitGroup // 等待组
)

func main() {
	RWMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	wg.Add(4)
	go writeData(1)
	go readData(2)
	go readData(3)
	go writeData(3)
	wg.Wait()
	fmt.Println("main over .......")
}

func writeData(i int) {
	defer wg.Done()

	fmt.Println(i, "开始写操作：write start ...")
	RWMutex.Lock() // 写操作上锁
	fmt.Println("正在写...")
	time.Sleep(3 * time.Second)
	RWMutex.Unlock() // 写解锁
	fmt.Println(i, "写结束:write over ...")
}

func readData(i int) {
	defer wg.Done()

	fmt.Println(i, "开始读：read start。。")
	RWMutex.RLock() // 读操作上锁
	fmt.Println(i, "正在读取数据：reading。。。")
	time.Sleep(3 * time.Second)
	RWMutex.RUnlock() //读操作解锁
	fmt.Println(i, "读结束：read over。。。")
}
