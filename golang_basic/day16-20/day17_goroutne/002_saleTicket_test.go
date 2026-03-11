package day17

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	ticket = 100          // 总票数
	wg     sync.WaitGroup // 等待组
	mu     sync.Mutex     // 互斥锁
)

// func TestSaleTicket(t *testing.T) {

// 	wg.Add(4) // 等待4个协程完成
// 	// 模拟四个售票窗口
// 	go saleTickets("窗口1")
// 	go saleTickets("窗口2")
// 	go saleTickets("窗口3")
// 	go saleTickets("窗口4")
// 	wg.Wait() // 等待所有协程完成
// 	// time.Sleep(10 * time.Second)
// }

// func saleTickets(name string) {
// 	defer wg.Done()
// 	// 定义一个随机数
// 	rand.Seed(time.Now().UnixNano())
// 	for {
// 		// 加锁
// 		mu.Lock()
// 		if ticket > 0 {
// 			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 随机休眠1000ms以内
// 			ticket--
// 			fmt.Println(name, "售出一张票，剩余", ticket, "张")
// 			// 解锁
// 			mu.Unlock()
// 		} else {
// 			fmt.Println(name, "售罄")
// 			// 解锁
// 			mu.Unlock()
// 			break
// 		}
// 	}
// }

// 出现数据竞争的原因是多个协程同时访问和修改共享变量ticket。为了解决这个问题，可以使用互斥锁来保护共享变量的访问。
// 以上会出现死锁

// 使用更加安全的锁管理

// func saleTickets(name string) {
// 	defer wg.Done()
// 	for {
// 		mu.Lock()

// 		// 检查票数
// 		if ticket <= 0 {
// 			mu.Unlock()
// 			fmt.Println("售罄")
// 			break
// 		}

// 		// 模拟售票时间
// 		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

// 		// 售票
// 		ticket--

// 		fmt.Println(name, "售出一张票，剩余", ticket, "张")

// 		// 解锁
// 		mu.Unlock()
// 	}
// }
// func TestSaleTicket2(t *testing.T) {
// 	wg.Add(4) // 等待4个协程完成
// 	// 模拟四个售票窗口
// 	go saleTickets("窗口1")
// 	go saleTickets("窗口2")
// 	go saleTickets("窗口3")
// 	go saleTickets("窗口4")
// 	wg.Wait() // 等待所有协程完成
// }

// 以上方案失败，主要是应为sleep总时间超过30s，测试用例默认不得超过30秒

// 优化方案
func TestSaleTicket3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	wg.Add(4)
	go saleTickets("窗口1")
	go saleTickets("窗口2")
	go saleTickets("窗口3")
	go saleTickets("窗口4")
	wg.Wait()
	fmt.Println("所有票已售完")
}

func saleTickets(name string) {
	defer wg.Done()
	for {
		mu.Lock()
		// 如果没票了，就退出循环
		if ticket <= 0 {
			mu.Unlock()
			fmt.Println("售罄")
			return
		}
		// 还有票，先减少票数
		ticket--
		currentTicket := ticket

		// 解锁
		mu.Unlock()

		// 模拟售票时间
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		fmt.Println(name, "售出一张票，剩余", currentTicket, "张")

		// 如果已经没票了就提前退出
		if currentTicket <= 0 {
			fmt.Println(name, "售罄")
			return
		}
	}
}
