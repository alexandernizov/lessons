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
	mx       sync.RWMutex
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func New(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
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

func (c *LRU) Get(key string) interface{} {
	c.mx.Lock()
	defer c.mx.Unlock()
	element, exists := c.items[key]
	if exists == false {
		return nil
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value
}
