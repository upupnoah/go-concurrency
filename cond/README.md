# cond

## Note

- cond 数据结构

  ```go
  type Cond struct {
    L Locker // 传入的锁（通常是一个 Mutex）
    notify  notifyList // 用于管理等待 goroutine 的队列
    checker copyChecker // 用于检测 Cond 是否被不安全地复制
  }
  ```

- cond 的方法

  - signal: 允许调用者 Caller 唤醒一个等待此 Cond 的 goroutine
  - broadcast: 允许调用者 Caller 唤醒所有等待此 Cond 的 goroutine
  - wait: 会把调用者 Caller 放入 Cond 的等待队列中并阻塞，直到被 Signal 或者 Broadcast 的方法从等待队列中移除并唤醒

- 使用场景

  - 等待某个条件成立
  - 通知其他等待的 goroutine

- 使用 cond 的两个常见错误
  - 调用 Wait 前忘记加锁
  - 只调用了一次 Wait，没有检查等待条件是否满足，结果条件没满足，程序就继续执行了
