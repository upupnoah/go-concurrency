package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	// mu    sync.Mutex
	sync.Mutex
	count int
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.count++
				counter.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.count)
}
