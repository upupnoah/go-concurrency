package main

type Cond struct {
	// noCopy noCopy

	// // 当观察或者修改等待条件的时候需要加锁
	// L Locker

	// // 等待队列
	// notify  notifyList
	// checker copyChecker
}

// func NewCond(l Locker) *Cond {
// 	return &Cond{L: l}
// }

func (c *Cond) Wait() {
	// c.checker.check()
	// 增加到等待队列中
	// t := runtime_notifyListAdd(&c.notify)
	// c.L.Unlock()
	// 阻塞休眠直到被唤醒
	// runtime_notifyListWait(&c.notify, t)
	// c.L.Lock()
}

func (c *Cond) Signal() {
	// c.checker.check()
	// runtime_notifyListNotifyOne(&c.notify)
}

func (c *Cond) Broadcast() {
	// c.checker.check()
	// runtime_notifyListNotifyAll(&c.notify）
}
