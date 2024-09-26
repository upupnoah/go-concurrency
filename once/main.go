package main

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	// 双检查
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
