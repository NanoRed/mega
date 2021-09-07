package memory

import (
	"time"

	"github.com/RedAFD/mega/internal/storage/memory/impl"
)

type memoryStorage interface {
	Get(key string) (value interface{}, exist bool)
	Set(key string, value interface{}, expiration ...time.Duration)
}

var memoryStorageEngine memoryStorage = impl.MemoryStorageGocache()

func Load(key string) (value interface{}, exist bool) {
	return memoryStorageEngine.Get(key)
}

func Store(key string, value interface{}, expiration ...time.Duration) {
	memoryStorageEngine.Set(key, value, expiration...)
}
