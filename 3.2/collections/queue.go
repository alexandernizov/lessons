package collections

type Queue struct {
	c Collection
}

func NewQueue() *Queue {
	c := New()
	return &Queue{c: *c}
}

func (q *Queue) Push(value any) {
	q.c.PushBack(value)
}

func (q *Queue) Pop() (any, bool) {
	return q.c.PopBack()
}

func (q *Queue) Peek() (any, bool) {
	return q.c.PeekBack()
}

func (q *Queue) Contains(value any) bool {
	return q.c.Contains(value)
}

func (q *Queue) Clear() {
	q.c.Clear()
}
