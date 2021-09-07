package impl

import (
	"sync"
	"time"

	"github.com/RedAFD/mega/internal/config"
	"github.com/patrickmn/go-cache"
)

var __memoryStorageGocache = &_memoryStorageGocache{}

type _memoryStorageGocache struct {
	cache *cache.Cache
	once  sync.Once
}

func (m *_memoryStorageGocache) Get(key string) (value interface{}, exist bool) {
	return m.cache.Get(key)
}

func (m *_memoryStorageGocache) Set(key string, value interface{}, expiration ...time.Duration) {
	if len(expiration) > 0 && expiration[0] > 0 {
		m.cache.Set(key, value, expiration[0])
		return
	}
	m.cache.Set(key, value, cache.NoExpiration)
}

func MemoryStorageGocache() *_memoryStorageGocache {
	__memoryStorageGocache.once.Do(func() {
		__memoryStorageGocache.cache = cache.New(config.GocacheDefaultExpiration, config.GocacheCleanupInterval)
	})
	return __memoryStorageGocache
}
