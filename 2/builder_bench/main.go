package main

import (
	"strconv"
	"strings"
)

func main() {
	//1. Проверить, что при конкатенации, у нас билдер работает быстрее.
	//	Ну как бы фиг знает, если просто складываем 3 строки - то не быстрее, а если 10 - то быстрее
	b := strings.Builder{}
	_ = b
	//2. Посмотреть как устроен билдер, и сравнить с п.1
	//3. Как работает разделение строк под капотом?
}

func Concat(strs ...string) string {
	return strings.Join(strs, " ")
}

func Build(strs ...string) string {
	var b strings.Builder
	b.Grow(len("0123456789") * 10)
	for _, str := range strs {
		b.WriteString(str)
	}
	return b.String()
}

func RuneToIntByString() {
	r := '9'
	strconv.Atoi(string(r))
}

func RuneToIntByMinus() {
	r := '9'
	i := r - '0'
	_ = int(i)
}
