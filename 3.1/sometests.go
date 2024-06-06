package main

import (
	"fmt"
)

func sometests() {
	array := [9]string{"a", "l", "e", "x", "a", "n", "d", "e", "r"}
	fmt.Println(array)
	//Массивы разной длины - это разные массивы
	arrayDifLen := [10]string{"", ""}
	_ = arrayDifLen
	// array2 = arrayDifLen // - не заработает, так как cannot use array (varablie of type [9]string as [10]string  value in)
	arraySameLen := [9]string{}
	fmt.Println(arraySameLen)
	arraySameLen = array      // - пустой
	fmt.Println(arraySameLen) // - копия
	fmt.Printf("%p\n", &array)
	fmt.Printf("%p\n", &arraySameLen) // у исходного и скопированного массивов разные адреса
	// Слайсы
	slice := []int{}
	fmt.Println(slice)
	sliceByMake := make([]int, 3, 5) // сразу указали длину и вместимость - элементы заполняются zero value
	fmt.Println(sliceByMake)
	emptySliceZeroValue := make([]MyStruct, 3, 5)
	fmt.Println(emptySliceZeroValue)
	emptySliceZeroValue2 := make([]bool, 3, 5)
	fmt.Println(emptySliceZeroValue2)
	// sliceWithLenMoreThanCap := make([]int, 5, 0) // будет ошибка на этапе компиляции)
	sliceSpreadingCapacity := make([]int, 0)
	v := 10
	_ = v
	fmt.Printf("%p\n", &sliceSpreadingCapacity)
	sliceSpreadingCapacity = append(sliceSpreadingCapacity, 2)
	fmt.Printf("%p\n", &sliceSpreadingCapacity)
	sliceSpreadingCapacity = append(sliceSpreadingCapacity, 3)
	fmt.Printf("%v\n", sliceSpreadingCapacity)
	//Пробуем создать массив, взять от него слайс, изменить элемент в массиве через слайс
	arr := [3]int{0, 0, 0} //Изначальный
	fmt.Println(arr)
	sl := arr[:]                         //Скопировали, не меняя длину
	sl[0] = 1                            //Поменяли первый элемент (по адресу)
	fmt.Println(arr)                     //Теперь первый элемент в оригинальном массиве стал 1
	longerSlice := arr[:]                //Создали новый массив
	longerSlice[0] = 2                   //Поменяли элемент в массиве
	fmt.Println(arr)                     //Теперь оба имеют 2
	fmt.Println(sl)                      //Теперь обма имеют 2 - ведь мы меняем по памяти
	longerSlice = append(longerSlice, 4) //Добавили новый элемент в слайс, и так как теперь уже не хватает длины - мы переопределили
	//слайс в новую память
	longerSlice[0] = 3       //поменяли элемент
	fmt.Println(longerSlice) //Смотрим, что новый слайс - изменился, а старые нет
	fmt.Println(arr)
	fmt.Println(sl)
	//Если нужно сделать копию массива и поменять что-то в копии, не меняя оригинальный слайс - используем копи
	copiedSl := make([]int, 2)
	fmt.Println("copy")
	copy(copiedSl, sl)
	copiedSl[1] = 1
	fmt.Println(copiedSl)
	fmt.Println(sl)
	//Но тогда нам надо знать длину слайса, иначе получится скопировать только то кол-во элементов, которое в нем было
	//А можем просто создать пустой слайс и сделать append
	appendedSl := make([]int, 0)
	appendedSl = append(appendedSl, sl...)
	fmt.Println(appendedSl)
	fmt.Println(sl)
	appendedSl[0] = 1
	fmt.Println(appendedSl)
	fmt.Println(sl)
	c := make(chan int)
	_ = c
	n := new(chan int)
	m := make(chan int)
	n = &m
	_ = n
}

type MyStruct struct {
}
