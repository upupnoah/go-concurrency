package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// RecursiveMutex 包装一个 Mutex, 实现可重入锁
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的 goroutine id
	recursion int32 // 这个 goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的 goroutine 就是这次调用的 goroutine, 说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的 goroutine 第一次调用，记录下它的 goroutine id, 调用次数加 1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	// 非持有锁的 goroutine 尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减 1
	m.recursion--
	if m.recursion != 0 { // 如果这个 goroutine 还没有完全释放，则直接返回
		return
	}
	// 此 goroutine 最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
