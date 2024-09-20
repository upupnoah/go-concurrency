# Mutex 互斥锁

## codes

- counter/main.go: 有并发安全的 counter, 因为 count++ 操作不是原子操作, 不是并发安全的
- counter-with-mutex/main.go: 使用 Mutex 来实现并发安全的 counter

## tips

- 可以使用 go run --race 来检查代码中的数据竞争问题
- sync.Mutex 零值可用
