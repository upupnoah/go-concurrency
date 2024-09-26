package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

func work(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go work(&counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.count)
}
