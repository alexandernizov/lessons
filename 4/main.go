package main

import (
	"fmt"

	"github.com/alexandenrizov/lessons/4/set"
)

func main() {
	// passws := lru.New(100)
	// for i := 0; i < 100; i++ {
	// 	go func(i int) {
	// 		passws.Set(strconv.Itoa(i), i)
	// 	}(i)
	// }
	// time.Sleep(1 * time.Second)
	// fmt.Println(passws.Get("99"))

	// passwsTTL := ttl.New(1 * time.Second)
	// for i := 0; i < 100; i++ {
	// 	go func(i int) {
	// 		passwsTTL.Set(strconv.Itoa(i), i)
	// 	}(i)
	// }
	// time.Sleep(500 * time.Millisecond)
	// passwsTTL.Set("1", 1)
	// time.Sleep(500 * time.Millisecond)
	// fmt.Println(passwsTTL.Get("1"))
	// fmt.Println(passwsTTL.Get("22"))
	s := set.New()
	s.Add(0)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s2 := set.New()
	s2.Add(0)
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)
	u := s.Union(s2)
	i := s.Intersection(s2)
	sub := s.Substraction(s2)
	sub2 := s2.Substraction(s)
	fmt.Println(u.M)
	fmt.Println(i.M)
	fmt.Println(sub.M)
	fmt.Println(sub2.M)
}
