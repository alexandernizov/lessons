package collections

type Stack struct {
	c Collection
}

func NewStack() *Stack {
	c := New()
	return &Stack{c: *c}
}

func (s *Stack) Push(value any) {
	s.c.PushBack(value)
}

func (s *Stack) Pop() (any, bool) {
	return s.c.PopBack()
}

func (s *Stack) Peek() (any, bool) {
	return s.c.PeekBack()
}
