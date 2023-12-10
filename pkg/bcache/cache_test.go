package bcache

import "testing"

func TestCache_GetString(t *testing.T) {
	cache := MemoryCache{}

	cache.GetString("key")
}
