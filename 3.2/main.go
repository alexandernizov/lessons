package main

import (
	"fmt"

	"github.com/alexandernizov/lessons/3.2/collections"
)

func main() {
	fmt.Println("Stack")
	st := collections.NewStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st.Peek())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())

	fmt.Println("Queue")
	q := collections.NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println("Deque")
	d := collections.NewDeque()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	fmt.Println(d.PeekBack())
	fmt.Println(d.PopBack())
	fmt.Println(d.PopBack())
	fmt.Println(d.PopBack())
	fmt.Println(d.PopBack())
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)
	fmt.Println(d.PeekFront())
	fmt.Println(d.PopFront())
	fmt.Println(d.PopFront())
	fmt.Println(d.PopFront())
	fmt.Println(d.PopFront())
}
