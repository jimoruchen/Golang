package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. 准备数据
	var data []string
	for i := 0; i < 100; i++ {
		data = append(data, fmt.Sprintf("Item-%d", i))
	}

	const workerCount = 10
	taskChan := make(chan int, 100) // 缓冲设为 100，防止发送阻塞，也可以设为 0
	var wg sync.WaitGroup

	// 2. 启动 10 个 Worker
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 从 channel 读取任务，直到 channel 关闭
			for idx := range taskChan {
				fmt.Printf("Worker %d processing: %s\n", id, data[idx])
			}
		}(i)
	}

	// 3. 主线程分发任务
	for i := 0; i < len(data); i++ {
		taskChan <- i // 发送索引
	}
	close(taskChan) // 任务发完，关闭 channel，通知 Worker 退出

	// 4. 等待结束
	wg.Wait()
	fmt.Println("All done!")
}
