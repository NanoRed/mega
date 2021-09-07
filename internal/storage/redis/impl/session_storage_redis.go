package impl

import (
	"context"
	"sync"
	"time"

	rds "github.com/RedAFD/mega/internal/storage/redis"
	"github.com/go-redis/redis/v8"
)

var __sessionStorageRedis = &_sessionStorageRedis{}

type _sessionStorageRedis struct {
	client *redis.Client
	once   sync.Once
}

func (s *_sessionStorageRedis) Set(key string, value []byte, expiration time.Duration) error {
	return s.client.SetEX(context.Background(), key, value, expiration).Err()
}

func (s *_sessionStorageRedis) Get(key string) (value []byte, err error) {
	strcmd := s.client.Get(context.Background(), key)
	if err = strcmd.Err(); err != nil {
		return
	}
	return strcmd.Bytes()
}

func SessionStorageRedis() *_sessionStorageRedis {
	__sessionStorageRedis.once.Do(func() {
		__sessionStorageRedis.client = rds.DB(rds.Session)
	})
	return __sessionStorageRedis
}
