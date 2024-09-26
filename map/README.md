# map

## 普通的 map

- 并发不安全
- 不保证顺序
  - 如果要保证顺序, 可以使用 [orderedmap](https://github.com/elliotchance/orderedmap)
  - 原理是设置数据结构的时候多封装了一个 list, 每次操作的时候都把 key 放到 list 的尾部
  - 要遍历顺序的时候, 直接遍历 list 即可

## 并发安全的 map

- 支持分片的 [concurrency-map](https://github.com/orcaman/concurrent-map)
- 使用 sync.Map
  - 比较好的使用场景
    - 只会增长的缓存系统中，一个 key 只写入一次而被读很多次
    - 多个 goroutine 为不相交的键集读、写和重写键值对

## sync.map 的实现

- 空间换时间。通过冗余的两个数据结构（只读的 read 字段、可写的 dirty），来减少加锁对性能的影响。对只读字段（read）的操作不需要加锁
- 优先从 read 字段读取、更新、删除，因为对 read 字段的读取不需要锁
- 动态调整。miss 次数多了之后，将 dirty 数据提升为 read，避免总是从 dirty 中加锁读取
- double-checking。加锁之后先还要再检查 read 字段，确定真的不存在才操作 dirty 字段
- 延迟删除。删除一个键值只是打标记，只有在提升 dirty 字段为 read 字段的时候才清理删除的数据

## Tips

![sync.map](https://static001.geekbang.org/resource/image/a8/03/a80408a137b13f934b0dd6f2b6c5cc03.jpg?wh=2250*1771)
