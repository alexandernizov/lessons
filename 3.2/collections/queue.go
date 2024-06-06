package collections

type Queue struct {
	c Collection
}

func NewQueue() *Queue {
	c := New()
	return &Queue{c: *c}
}

func (s *Queue) Push(value any) {
	s.c.PushBack(value)
}

func (s *Queue) Pop() (any, bool) {
	return s.c.PopBack()
}

func (s *Queue) Peek() (any, bool) {
	return s.c.PeekBack()
}
