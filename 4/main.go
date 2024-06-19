package main

import (
	"fmt"

	"github.com/alexandenrizov/lessons/4/cache"
	"github.com/alexandenrizov/lessons/4/cache/lru"
	"github.com/alexandenrizov/lessons/4/cache/ttl"
)

func main() {
	cch := cache.New(cache.SetTTL(ttl.New()))
	cch.Set("1", 1)
	cch.Set("2", 2)
	cch.Set("3", 2)
	fmt.Println(cch.Get("4"))

	cch2 := cache.New(cache.SetLRU(lru.New(lru.SetCapacity(2))))
	cch2.Set("1", "qwe")
	cch2.Set("2", "asd")
	cch2.Set("3", "zxc")
	fmt.Println(cch2.Get("3"))
}
