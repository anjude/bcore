package bcache

import (
	"github.com/patrickmn/go-cache"
)

type MemoryCache cache.Cache

func (c *MemoryCache) GetString(key string) (string, bool) {
	if v, ok := c.Get(key); ok {
		return v.(string), ok
	}
	return "", false
}

func (c *MemoryCache) GetInt64(key string) (int64, bool) {
	if v, ok := c.Get(key); ok {
		return v.(int64), ok
	}
	return 0, false
}

func (c *MemoryCache) GetBool(key string) (bool, bool) {
	if v, ok := c.Get(key); ok {
		return v.(bool), ok
	}
	return false, false
}
