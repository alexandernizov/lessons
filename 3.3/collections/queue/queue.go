package queue

import "github.com/alexandernizov/lessons/3.3/collections"

type Queue struct {
	c Collection
}

func New() *Queue {
	c := collections.New()
	return &Queue{c}
}

type Collection interface {
	PushBack(value any)
	PopFront() (any, bool)
	PeekFront() (any, bool)
	Contains(value any) bool
	Clear()
}

func (q *Queue) Push(value any) {
	q.c.PushBack(value)
}

func (q *Queue) Pop() (any, bool) {
	return q.c.PopFront()
}

func (q *Queue) Peek() (any, bool) {
	return q.c.PeekFront()
}

func (q *Queue) Contains(value any) bool {
	return q.c.Contains(value)
}

func (q *Queue) Clear() {
	q.c.Clear()
}
