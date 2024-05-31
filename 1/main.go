package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	aaa [2]bool // 2 байта
	ccc [2]bool // 2 байта
	bbb int32   // 4 байта
}

func main() {
	var s Foo
	x := &s
	y := s

	fmt.Println(unsafe.Sizeof(x)) // 8
	fmt.Println(unsafe.Sizeof(y)) // 8
}
