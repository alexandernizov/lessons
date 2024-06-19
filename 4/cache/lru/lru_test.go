package lru_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/alexandenrizov/lessons/4/cache/lru"
	"github.com/stretchr/testify/assert"
)

func TestLRUmutex(t *testing.T) {
	c := lru.New(lru.SetCapacity(100))
	for i := 0; i < 100; i++ {
		go func(i int) {
			c.Set(strconv.Itoa(i), i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	v, _ := c.Get("99")
	assert.Equal(t, 99, v, "Incorrect work")
}

func TestLRUoverflow(t *testing.T) {
	c := lru.New(lru.SetCapacity(2))
	c.Set("1", 1)
	c.Set("2", 2)
	c.Set("1", 1)
	c.Set("3", 3)
	v1, _ := c.Get("1")
	v2, _ := c.Get("2")
	v3, _ := c.Get("3")
	assert.Equal(t, 1, v1, "Incorrect work")
	assert.Equal(t, nil, v2, "Incorrect work")
	assert.Equal(t, 3, v3, "Incorrect work")
}
