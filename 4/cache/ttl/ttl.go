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
	mx            *sync.RWMutex
	ttl           time.Duration
	cleanInterval time.Duration
	items         map[string]Item
}

func New(opts ...ttlOption) *TTL {
	//Создали кэш ТТЛ
	defaultTTL := 1 * time.Second
	defaultCleanInterval := 500 * time.Millisecond
	defaultMutex := sync.RWMutex{}

	cache := TTL{
		mx:            &defaultMutex,
		ttl:           defaultTTL,
		cleanInterval: defaultCleanInterval,
		items:         make(map[string]Item),
	}
	for _, opt := range opts {
		opt(&cache)
	}

	go cache.garbageCollector()

	return &cache
}

type ttlOption func(*TTL)

func SetTTL(ttl time.Duration) ttlOption {
	return func(c *TTL) {
		c.ttl = ttl
	}
}

func SetClearInterval(ci time.Duration) ttlOption {
	return func(c *TTL) {
		c.cleanInterval = ci
	}
}

func SetMutex(mx *sync.RWMutex) ttlOption {
	return func(c *TTL) {
		c.mx = mx
	}
}

func (c *TTL) garbageCollector() {
	for {
		<-time.After(c.cleanInterval)
		c.clearExpired()
	}
}

func (c *TTL) Set(key string, value any) bool {
	return c.SetIndividual(key, value, 0)
}

func (c *TTL) SetIndividual(key string, value any, ttl time.Duration) bool {
	c.mx.Lock()
	defer c.mx.Unlock()

	var dur time.Duration
	if ttl == 0 {
		dur = c.ttl
	}

	//Здесь уже проверки на существование - не нужно, можем просто обновлять данные по ключу, даже если он существует
	c.items[key] = Item{
		Expiration: time.Now().Add(dur),
		Value:      value,
	}
	return true
}

func (c *TTL) Get(key string) (any, bool) {
	//Возьмем значение из кэша, если его нет - возвращаем nil, false
	//Если нашли значение, но его срок действия истек - возвращаем nil, false
	//Если нашли значение, и его срок действия не истек - возвращаем значение, true
	c.mx.RLock()
	defer c.mx.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if item.Expiration.Before(time.Now()) {
		return nil, false
	}

	return item.Value, true
}

func (c *TTL) clearExpired() {
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
