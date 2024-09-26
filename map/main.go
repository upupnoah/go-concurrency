package main

import "sync"

func main() {
	m := sync.Map{}
	m.Store(1, 1)
	m.Load(1)
}
