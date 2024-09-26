# Context

## 什么是 Context

- Context 是 Go 语言中用于在多个 goroutine 之间传递请求范围的值、取消信号和其他信息的结构体
- Context 是线程安全的，多个 goroutine 可以并发地使用同一个 Context
- Context 是链式结构，可以包含多个 Context，形成一个 Context 链

## Context 接口

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

- 关于 Done 方法: 如果 Done 没有被 close，Err 方法返回 nil；如果 Done 被 close，Err 方法会返回 Done 被 close 的原因

## context 包提供的方法

- context.Background(): 返回一个空的 Context，通常用于根 Context
- context.TODO(): 返回一个空的 Context，通常用于不确定使用哪个 Context 的情况
- context.WithCancel(): 返回一个新的 Context，当调用 cancel 函数时，会向 Context 发送取消信号
- context.WithTimeout(): 返回一个新的 Context，当到达指定时间时，会向 Context 发送取消信号
- context.WithDeadline(): 返回一个新的 Context，当到达指定时间时，会向 Context 发送取消信号
- context.WithValue(): 返回一个新的 Context，并携带一个键值对

## Context 的用途

- 传递请求范围的值
- 传递取消信号
- 传递截止时间
- 传递其他信息

## Context 的实现

- Context 的实现主要依赖于两个结构体：
  - context.Context
  - context.cancelCtx

## 使用 context 时约定俗成的规则

- 一般函数使用 Context 的时候，会把这个参数放在第一个参数的位置
- 从来不把 nil 当做 Context 类型的参数值，可以使用 context.Background() 创建一个空的上下文对象，也不要使用 nil
- Context 只用来临时做函数之间的上下文透传，不能持久化 Context 或者把 Context 长久保存。把 Context 持久化到数据库、本地文件或者全局变量、缓存中都是错误的用法
- key 的类型不应该是字符串类型或者其它内建类型，否则容易在包之间使用 Context 时候产生冲突。使用 WithValue 时，key 的类型应该是自己定义的类型
- 常常使用 struct{}作为底层类型定义 key 的类型。对于 exported key 的静态类型，常常是接口或者指针。这样可以尽量减少内存分配

## Tips

![context](https://static001.geekbang.org/resource/image/2d/2b/2dcbb1ca54c31b4f3e987b602a38e82b.jpg?wh=2250*2441)

## Note

- 可以参考飞雪无情的 [context 笔记](https://www.flysnow.org/2017/05/12/go-in-action-go-context)