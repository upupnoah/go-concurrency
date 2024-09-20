package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0

	var wg sync.WaitGroup
	wg.Add(10)

	var mu sync.Mutex // 零值可用

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
