package collections

type Dueue struct {
	c Collection
}

func NewDeque() *Dueue {
	c := New()
	return &Dueue{c: *c}
}

func (s *Dueue) PushBack(value any) {
	s.c.PushBack(value)
}

func (s *Dueue) PopBack() (any, bool) {
	return s.c.PopBack()
}

func (s *Dueue) PeekBack() (any, bool) {
	return s.c.PeekBack()
}

func (s *Dueue) PushFront(value any) {
	s.c.PushFront(value)
}

func (s *Dueue) PopFront() (any, bool) {
	return s.c.PopFront()
}

func (s *Dueue) PeekFront() (any, bool) {
	return s.c.PeekFront()
}
