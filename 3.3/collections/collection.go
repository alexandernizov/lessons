package collections

import (
	"sync"
)

type Collection struct {
	new   bool
	mx    *sync.RWMutex
	slice []any
}

func New() *Collection {
	mx := sync.RWMutex{}
	return &Collection{new: true, mx: &mx}
}

func (c *Collection) wasnew() {
	if !c.new {
		panic("was initialized without New()")
	}
}

func (c *Collection) PushBack(value any) {
	c.wasnew()
	c.mx.Lock()
	defer c.mx.Unlock()
	c.slice = append(c.slice, value)
}

func (c *Collection) PopBack() (any, bool) {
	c.wasnew()
	c.mx.Lock()
	defer c.mx.Unlock()
	if len(c.slice) <= 0 {
		return nil, false
	}
	res := c.slice[len(c.slice)-1]
	c.slice = c.slice[:len(c.slice)-1]
	return res, true
}

func (c *Collection) PeekBack() (any, bool) {
	c.wasnew()
	c.mx.RLock()
	defer c.mx.RUnlock()
	if len(c.slice) <= 0 {
		return nil, false
	}
	res := c.slice[len(c.slice)-1]
	return res, true
}

func (c *Collection) PushFront(value any) {
	c.wasnew()
	c.mx.Lock()
	defer c.mx.Unlock()
	c.slice = append(c.slice, value)
	copy(c.slice[1:], c.slice)
	c.slice[0] = value
}

func (c *Collection) PopFront() (any, bool) {
	c.wasnew()
	c.mx.Lock()
	defer c.mx.Unlock()
	if len(c.slice) <= 0 {
		return "", false
	}
	res := c.slice[0]
	c.slice = c.slice[1:]
	return res, true
}

func (c *Collection) PeekFront() (any, bool) {
	c.wasnew()
	c.mx.RLock()
	defer c.mx.RUnlock()
	if len(c.slice) <= 0 {
		return "", false
	}
	res := c.slice[0]
	return res, true
}

func (c *Collection) Contains(value any) bool {
	c.wasnew()
	c.mx.RLock()
	defer c.mx.RUnlock()
	if len(c.slice) <= 0 {
		return false
	}
	for _, v := range c.slice {
		if v == value {
			return true
		}
	}
	return false
}

func (c *Collection) Clear() {
	c.wasnew()
	c.mx.Lock()
	defer c.mx.Unlock()
	c.slice = nil
}
