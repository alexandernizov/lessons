package lru_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/alexandenrizov/lessons/4/cache/lru"
	"github.com/stretchr/testify/assert"
)

func TestLRUmutex(t *testing.T) {
	c := lru.New(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			c.Set(strconv.Itoa(i), i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 99, c.Get("99"), "Incorrect work")
}

func TestLRUoverflow(t *testing.T) {
	c := lru.New(2)
	c.Set("1", 1)
	c.Set("2", 2)
	c.Set("1", 1)
	c.Set("3", 3)
	assert.Equal(t, 1, c.Get("1"), "Incorrect work")
	assert.Equal(t, nil, c.Get("2"), "Incorrect work")
	assert.Equal(t, 3, c.Get("3"), "Incorrect work")
}
