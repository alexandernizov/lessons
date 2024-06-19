package cache

import (
	"github.com/alexandenrizov/lessons/4/cache/lru"
	"github.com/alexandenrizov/lessons/4/cache/ttl"
)

type Cache struct {
	c Cachier
}

type Cachier interface {
	Set(key string, value any) bool
	Get(key string) (any, bool)
}

type cacheOptions func(*Cache)

func New(opts ...cacheOptions) *Cache {
	defaultCache := lru.New()

	cache := Cache{c: defaultCache}

	for _, opt := range opts {
		opt(&cache)
	}

	return &cache
}

func SetLRU(lru *lru.LRU) cacheOptions {
	return func(c *Cache) {
		c.c = lru
	}
}

func SetTTL(ttl *ttl.TTL) cacheOptions {
	return func(c *Cache) {
		c.c = ttl
	}
}

func (c *Cache) Set(key string, value any) {
	c.c.Set(key, value)
}

func (c *Cache) Get(key string) (any, bool) {
	return c.c.Get(key)
}
