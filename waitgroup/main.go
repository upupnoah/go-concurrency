package main

import (
	"sync/atomic"
	"unsafe"
)

type WaitGroup struct {
	// 避免复制使用的一个技巧，可以告诉vet工具违反了复制使用的规则
	// noCopy noCopy
	// 64bit(8bytes)的值分成两段，高32bit是计数值，低32bit是waiter的计数
	// 另外32bit是用作信号量的
	// 因为64bit值的原子操作需要64bit对齐，但是32bit编译器不支持，所以数组中的元素在不同的架构中不一样，具体处理看下面的方法
	// 总之，会找到对齐的那64bit作为state，其余的32bit做信号量
	state1 [3]uint32
}

// 得到state的地址和信号量的地址
func (wg *WaitGroup) state() (statep *uint64, semap *uint32) {
	if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
		// 如果地址是64bit对齐的，数组前两个元素做state，后一个元素做信号量
		return (*uint64)(unsafe.Pointer(&wg.state1)), &wg.state1[2]
	} else {
		// 如果地址是32bit对齐的，数组后两个元素用来做state，它可以用来做64bit的原子操作，第一个元素32bit用来做信号量
		return (*uint64)(unsafe.Pointer(&wg.state1[1])), &wg.state1[0]
	}
}

func (wg *WaitGroup) Add(delta int) {
	statep, _ := wg.state()
	// 高32bit是计数值v，所以把delta左移32，增加到计数上
	state := atomic.AddUint64(statep, uint64(delta)<<32)
	v := int32(state >> 32) // 当前计数值
	w := uint32(state)      // waiter count

	if v > 0 || w == 0 {
		return
	}

	// 如果计数值v为0并且waiter的数量w不为0，那么state的值就是waiter的数量
	// 将waiter的数量设置为0，因为计数值v也是0,所以它们俩的组合*statep直接设置为0即可。此时需要并唤醒所有的waiter
	*statep = 0
	for ; w != 0; w-- {
		// runtime_Semrelease(semap, false, 0)
	}
}

// Done方法实际就是计数器减1
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}
