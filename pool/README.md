# pool (对象池)

## sync.Pool

- 使用 sync.Pool 的场景

  - 一个对象的创建成本比较高
  - 一个对象的复用周期比较长
  - 一个对象的复用频率比较高
  - 一个对象的复用场景比较独立，不希望和其他的实例混用

- 需要知道的内容
  - sync.Pool 本身就是线程安全的，多个 goroutine 可以并发地调用它的方法存取对象
  - sync.Pool 不可在使用之后再复制使用

## sync.Pool 数据结构

![sync.Pool 数据结构](https://static001.geekbang.org/resource/image/f4/96/f4003704663ea081230760098f8af696.jpg?wh=3659*2186)
