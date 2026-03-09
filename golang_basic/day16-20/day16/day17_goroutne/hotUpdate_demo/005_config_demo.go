package main

import (
	"fmt"
	"sync"
)

/**
	实现目标：
	1.集中管理配置：统一存储应用的各项配置
	2.线程安全读取：多个goroutine可以并发读取配置
	3.动态热更新：运行是可以更新配置，无需重启
	4.读写分离：读操作多，写操作少
	5.批量操作：支持批量更新多个配置项

**/

type ConfigManager struct {
	config map[string]interface{} //
	rwmu   sync.RWMutex
}

// 读操作
func (cm *ConfigManager) GetAllConfig(key string) map[string]interface{} {
	cm.rwmu.RLock()
	defer cm.rwmu.RUnlock()

	result := make(map[string]interface{})
	for k, v := range result {
		result[k] = v
	}
	return result
	/*
		返回深拷贝而不是原始map → 防止外部代码意外修改内部配置
		虽然增加了一点内存开销，但保证了数据安全*/
}

// 更新单个操作
func (cm *ConfigManager) UpdateConfig(key string, value interface{}) {
	cm.rwmu.Lock()
	defer cm.rwmu.Unlock()

	cm.config[key] = value
	fmt.Printf("配置更新：%s=%v \n", key, value)
}

// 批量更新
func (cm *ConfigManager) BatchUpdateConfigs(updates map[string]interface{}) {
	cm.rwmu.Lock()
	defer cm.rwmu.Unlock()

	for k, v := range updates {
		cm.config[k] = v
		fmt.Printf("批量更新：%s=%v \n", k, v)
	}
}

// 场景：配置热更新时
/*
时间线：
1. 10个goroutine正在并发读取配置（都持有读锁）
2. 管理后台请求更新配置
3. 更新操作尝试获取写锁，但被阻塞（因为有读锁存在）
4. 已有读锁全部释放后，更新操作获得写锁
5. 更新配置期间，新来的读操作被阻塞（因为写锁存在）
6. 更新完成，释放写锁
7. 等待的读操作获得读锁，读取到新配置
*/
// func main() {
// 	configMgr := NewConfigManager()
// 	var wg sync.WaitGroup

// 	// 模拟5个服务并发读取配置
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go func(serviceID int) {
// 			defer wg.Done()
// 			for j := 0; j < 3; j++ {
// 				time.Sleep(time.Duration(serviceID*100) * time.Millisecond)
// 				if value, ok := configMgr.GetConfig("max_workers"); ok {
// 					fmt.Printf("服务%d读取配置: max_workers = %v\n",
// 						serviceID, value)
// 				}
// 			}
// 		}(i)
// 	}

// 	// 模拟配置热更新（在后台进行）
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		time.Sleep(200 * time.Millisecond)        // 等待一会儿
// 		configMgr.UpdateConfig("max_workers", 20) // 更新配置
// 	}()

// 	wg.Wait()
// }
