package main

import (
	"fmt"

	"github.com/alexandernizov/lessons/3.3/collections/stack"
)

func main() {
	fmt.Println("Stack")

	st := stack.New()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
}
