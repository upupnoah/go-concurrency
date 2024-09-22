package main

// 伪代码 (pseudo code)

// CAS操作，当时还没有抽象出atomic包
func cas(val *int32, old, new int32) bool {
	return true
}

func semacquire(*int32) {}
func semrelease(*int32) {}

// 互斥锁的结构，包含两个字段
type Mutex struct {
	key  int32 // 锁是否被持有的标识
	sema int32 // 信号量专用，用以阻塞/唤醒goroutine
}

// 保证成功在val上增加delta的值
func xadd(val *int32, delta int32) (new int32) {
	for {
		v := *val
		if cas(val, v, v+delta) {
			return v + delta
		}
	}
	panic("unreached")
}

// Noah 的理解
// 1. 对于加锁, 在执行 xadd 1 之后, 如果 == 1, 则表示加锁成功, 否则阻塞等待
// 2. 对于解锁, 在执行 xadd -1 之后, 如果 == 0, 则表示解锁成功, 否则唤醒等待者

// 请求锁
func (m *Mutex) Lock() {
	// 原子操作的必要性:
	// 1. 如果多个goroutine同时对一个变量进行操作, 可能会导致数据不一致的问题
	// 2. 通过原子操作, 可以保证对变量的操作是原子的, 即在操作完成之前, 不会有其它goroutine对变量进行操作
	if xadd(&m.key, 1) == 1 { // 标识加1，如果等于1，成功获取到锁
		return
	}
	semacquire(&m.sema) // 否则阻塞等待
}

// 释放锁
func (m *Mutex) Unlock() {
	if xadd(&m.key, -1) == 0 { // 将标识减去1，如果等于0，则没有其它等待者
		return
	}
	semrelease(&m.sema) // 唤醒其它阻塞的 goroutine
}
