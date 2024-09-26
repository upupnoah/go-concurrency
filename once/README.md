# once

## 使用场景

- Once 可以用来执行且仅仅执行一次动作，常常用于单例对象的初始化场景
- 在多线程环境下，可以保证某个动作只执行一次
- 一旦你遇到只需要初始化一次的场景，首先想到的就应该是 Once 并发原语

## 使用 Once 可能出现的两种错误

1. 死锁
2. 未初始化成功
   - 解决方法: 自己实现一个 Once, 在 do 的时候返回 error

## Tips

![once](https://static001.geekbang.org/resource/image/4b/ba/4b1721a63d7bd3f3995eb18cee418fba.jpg?wh=2250*880)
