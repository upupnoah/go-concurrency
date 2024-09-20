package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0

	var wg sync.WaitGroup
	wg.Add(10) // 使用 WaitGroup 等待 10 个 goroutine 完成

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++ // 这不是一个原子操作, 不是并发安全的, 因此不会产生自己期望的操作
			}
		}()
	}

	wg.Wait() // 等待 10 个 goroutine 完成
	fmt.Println(count)
}
