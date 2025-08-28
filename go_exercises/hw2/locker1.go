package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter := SafeCounter{}

	// 启动10个goroutine同时增加计数
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second)

	// 输出最终计数
	fmt.Printf("Final count: %d\n", counter.GetCount())
}
