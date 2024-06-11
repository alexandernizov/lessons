package ttl_test

import (
	"testing"
	"time"

	"github.com/alexandenrizov/lessons/4/cache/ttl"
	"github.com/stretchr/testify/assert"
)

func TestTTL(t *testing.T) {
	c := ttl.New(100 * time.Millisecond)
	c.Set("1", 1)
	time.Sleep(50 * time.Millisecond)
	c.Set("2", 2)
	time.Sleep(50 * time.Millisecond)

	val, exists := c.Get("1")
	assert.Equal(t, nil, val, "Incorrect result")
	assert.Equal(t, false, exists, "Incorrect result")
	val, exists = c.Get("2")
	assert.Equal(t, 2, val, "Incorrect result")
	assert.Equal(t, true, exists, "Incorrect result")
	val, exists = c.Get("3")
	assert.Equal(t, nil, val, "Incorrect result")
	assert.Equal(t, false, exists, "Incorrect result")
}

func TestTTLClear(t *testing.T) {
	c := ttl.New(1 * time.Millisecond)
	c.ClearExpired()
	c.Set("1", 1)
	c.Set("2", 2)
	c.Set("3", 3)
	time.Sleep(2 * time.Millisecond)
	c.ClearExpired()
	assert.Equal(t, 0, c.Len(), "Incorrect result")
}
