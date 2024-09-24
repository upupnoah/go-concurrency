package main

import "sync"

type SliceQueue struct {
	data []any
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]any, 0, n)}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v any) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, v)
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() any {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) == 0 {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
