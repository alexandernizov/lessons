package ttl

import (
	"sync"
	"time"
)

type Item struct {
	Expiration time.Time
	Value      any
}

type TTL struct {
	mx    sync.RWMutex
	ttl   time.Duration
	items map[string]Item
}

func New(ttl time.Duration) *TTL {
	return &TTL{
		ttl:   ttl,
		items: make(map[string]Item),
	}
}

func (c *TTL) Set(key string, value any) {
	c.mx.Lock()
	defer c.mx.Unlock()

	//Здесь уже проверки на существование - не нужно, можем просто обновлять данные по ключу, даже если он существует
	c.items[key] = Item{
		Expiration: time.Now().Add(c.ttl),
		Value:      value,
	}
}

func (c *TTL) Get(key string) (any, bool) {
	//Возьмем значение из кэша, если его нет - возвращаем nil, false
	//Если нашли значение, но его срок действия истек - удаляем его, возвращаем nil, false
	//Если нашли значение, и его срок действия не истек - возвращаем значение, true
	c.mx.Lock()
	defer c.mx.Unlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if item.Expiration.Before(time.Now()) {
		delete(c.items, key)
		return nil, false
	}

	return item.Value, true
}

func (c *TTL) ClearExpired() {
	keys := []string{}
	t := time.Now()
	c.mx.Lock()
	for k, v := range c.items {
		if v.Expiration.Before(t) {
			keys = append(keys, k)
		}
	}
	if len(keys) == 0 {
		c.mx.Unlock()
		return
	}
	for _, v := range keys {
		delete(c.items, v)
	}
	c.mx.Unlock()
}

func (c *TTL) Len() int {
	return len(c.items)
}
