package main

import (
	"fmt"
	"time"
)

// https://cuihairu.github.io/hello-golang/concurrency/CSP.html
// CPS 通信顺序进程 将系统的并发执行单位抽象为进程 进程之间通过通道来传递消息

// 模拟一个并发任务

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d started job %d \n", id, j)

		time.Sleep(time.Second)

		fmt.Printf("worker %d finished job %d \n", id, j)

		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// 启动三个协程
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送五个任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for a := 1; a <= numJobs; a++ {
		<-results
		// fmt.Println("result:", <-results)
	}

}

// worker 3 started job 3
// worker 1 started job 1
// worker 2 started job 2
// worker 2 finished job 2
// worker 2 started job 4
// worker 1 finished job 1
// worker 1 started job 5
// worker 3 finished job 3
// worker 1 finished job 5
// worker 2 finished job 4

//   执行顺序：
//   1. main创建两个通道：jobs和results
//   2. main启动3个worker goroutine
//   3. main向jobs通道发送5个任务（1,2,3,4,5）
//   4. 3个worker同时从jobs通道抢夺任务
//   5. 每个worker处理完任务后，将结果发送到results通道
//   6. main从results通道接收5次结果并打印

// worker抢夺是不确定的直至五个任务全部完成

// ● 3个worker从jobs通道"抢夺"任务的机制是这样的：

//   Channel的内部机制：
//   - jobs通道是一个FIFO队列（先进先出）
//   - 当多个goroutine同时从同一个通道读取时，Go运行时会自动调度
//   - 每个任务只会被一个worker接收到

//   具体抢夺过程：

//   1. 任务分配：
//   jobs通道: [1, 2, 3, 4, 5]
//   worker1, worker2, worker3 同时执行 for j := range jobs
//   2. 可能的分配结果：
//   worker1 可能抢到: 任务1, 任务4
//   worker2 可能抢到: 任务2, 任务5
//   worker3 可能抢到: 任务3
//   3. 时间序列：
//   t=0s: worker1抢到任务1, worker2抢到任务2, worker3抢到任务3
//   t=1s: 前3个任务完成，worker1抢到任务4, worker2抢到任务5
//   t=2s: 所有任务完成

//   关键特性：
//   - 原子性：每个任务只会被一个worker接收
//   - 公平性：Go运行时尽量公平分配，但不保证绝对平均
//   - 并发安全：不会出现两个worker同时处理同一个任务
//   - 阻塞机制：如果没有任务，worker会阻塞等待
