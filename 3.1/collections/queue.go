package collections

import "sync"

type Queue struct {
	new   bool
	mx    *sync.RWMutex
	slice []any
}

func (q *Queue) wasnew() {
	if !q.new {
		panic("was initialized without NewQueue()")
	}
}

func NewQueue() *Queue {
	mx := sync.RWMutex{}
	return &Queue{new: true, mx: &mx}
}

func (q *Queue) Clear() {
	q.wasnew()
	q.mx.Lock()
	defer q.mx.Unlock()
	q.slice = nil
}

func (q *Queue) Contains(value any) bool {
	q.wasnew()
	q.mx.RLock()
	defer q.mx.Unlock()
	if len(q.slice) <= 0 {
		return false
	}
	for _, v := range q.slice {
		if v == value {
			return true
		}
	}
	return false
}

func (q *Queue) Dequeue() any {
	q.wasnew()
	q.mx.Lock()
	defer q.mx.Unlock()
	if len(q.slice) <= 0 {
		return nil
	}
	res := q.slice[0]
	q.slice = q.slice[1:]
	return res
}

func (q *Queue) Enqueue(value any) {
	q.wasnew()
	q.mx.Lock()
	defer q.mx.Unlock()
	q.slice = append(q.slice, value)
}

func (q *Queue) Peek() any {
	q.wasnew()
	q.mx.RLock()
	defer q.mx.RUnlock()
	if len(q.slice) <= 0 {
		return nil
	}
	res := q.slice[0]
	return res
}
