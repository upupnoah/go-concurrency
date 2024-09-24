package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func try() {
	var mu Mutex
	go func() { // 启动一个goroutine持有一段时间的锁
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock() // 尝试获取到锁
	if ok {            // 获取成功
		fmt.Println("got the lock")
		// do something
		mu.Unlock()
		return
	}

	// 没有获取到
	fmt.Println("can't get the lock")
}

func testCount() {
	var mu Mutex
	for i := 0; i < 1000; i++ { // 启动 1000 个 goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	// 输出锁的信息
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t,  starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}

func testQueue() {
	q := NewSliceQueue(10)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

func main() {
	try()
	testCount()
	testQueue()
}
