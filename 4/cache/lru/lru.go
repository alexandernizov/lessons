package lru

import (
	"container/list"
	"sync"
)

type Item struct {
	Key   string
	Value any
}

type LRU struct {
	mx       *sync.Mutex
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func New(opts ...lruOption) *LRU {
	defaultCapacity := 100 //Если оставить defaultCapcity == 0 - то тогда кэш будет неограниченного размера, и как его чистить?
	defaultMutex := sync.Mutex{}

	cache := LRU{
		mx:       &defaultMutex,
		capacity: defaultCapacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
	for _, opt := range opts {
		opt(&cache)
	}

	return &cache
}

type lruOption func(*LRU)

func SetCapacity(capacity int) lruOption {
	return func(c *LRU) {
		c.capacity = capacity
	}
}

func SetMutex(mx *sync.Mutex) lruOption {
	return func(c *LRU) {
		c.mx = mx
	}
}

func (c *LRU) Set(key string, value any) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	//Если по данному ключу уже существует запись в очереди - то мы двигаем запрашиваемый элемент в начало очереди
	element, exists := c.items[key]

	if exists {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	//Если по ключу не существует записи в списке - то надо его записать, а сначала смотрим, есть ли у нас для этого место
	if c.queue.Len() == c.capacity {
		c.purge()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element = c.queue.PushFront(item)
	c.items[item.Key] = element
	return true
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}

func (c *LRU) Get(key string) (any, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	element, exists := c.items[key]
	if !exists {
		return nil, false
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value, true
}
