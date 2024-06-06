package dueue

import "github.com/alexandernizov/lessons/3.3/collections"

type Dueue struct {
	c Collection
}

func New() *Dueue {
	c := collections.New()
	return &Dueue{c}
}

type Collection interface {
	PushBack(value any)
	PopBack() (any, bool)
	PeekBack() (any, bool)
	PushFront(value any)
	PopFront() (any, bool)
	PeekFront() (any, bool)
	Contains(value any) bool
	Clear()
}

func (d *Dueue) PushBack(value any) {
	d.c.PushBack(value)
}

func (d *Dueue) PopBack() (any, bool) {
	return d.c.PopBack()
}

func (d *Dueue) PeekBack() (any, bool) {
	return d.c.PeekBack()
}

func (d *Dueue) PushFront(value any) {
	d.c.PushFront(value)
}

func (d *Dueue) PopFront() (any, bool) {
	return d.c.PopFront()
}

func (d *Dueue) PeekFront() (any, bool) {
	return d.c.PeekFront()
}

func (d *Dueue) Contains(value any) bool {
	return d.c.Contains(value)
}

func (d *Dueue) Clear() {
	d.c.Clear()
}
