package main

import (
	"log"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		// 在 Go 新版本之后, 可以不传入外部的值到 goroutine 中
		go func() {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}()
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	//所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}
