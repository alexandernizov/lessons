package stack

import "github.com/alexandernizov/lessons/3.3/collections"

type Stack struct {
	c Collection
}

func New() *Stack {
	c := collections.New()
	return &Stack{c}
}

type Collection interface {
	PushBack(value any)
	PopBack() (any, bool)
	PeekBack() (any, bool)
	Contains(value any) bool
	Clear()
}

func (s *Stack) Push(value any) {
	s.c.PushBack(value)
}

// Returning value and true if the stack has a value, or false if stack doesn't have any value
func (s *Stack) Pop() (any, bool) {
	return s.c.PopBack()
}

func (s *Stack) Peek() (any, bool) {
	return s.c.PeekBack()
}

func (s *Stack) Contains(value any) bool {
	return s.c.Contains(value)
}

func (s *Stack) Clear() {
	s.c.Clear()
}
