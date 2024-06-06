package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	i, _ := strconv.Atoi("10")
	_ = i
	fmt.Println("asd")
	b := strings.Builder{}
	_ = b
	unicode.IsDigit('9')
}
