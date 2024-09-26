package main

import (
	"bytes"
	"fmt"
	"sync"
)

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		// 每次都创建一个新的 bytes.Buffer
// 		// buf := new(bytes.Buffer)
// 		buf := bytes.Buffer{}
// 		buf.WriteString("Hello, World!")
// 		fmt.Println(buf.String())
// 	}
// }

// 使用 sync.Pool
func main() {
	pool := sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	// 使用 sync.Pool 的好处
	// 1. 减少内存分配
	// 2. 减少 GC 压力
	for i := 0; i < 1000; i++ {
		buf := pool.Get().(*bytes.Buffer)
		buf.Reset() // 重置缓冲区
		buf.WriteString("Hello, World!")
		fmt.Println(buf.String())
		pool.Put(buf)
	}
}
