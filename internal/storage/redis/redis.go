package redis

import (
	"context"
	"sync"

	"github.com/RedAFD/mega/internal/config"
	"github.com/RedAFD/mega/internal/utils/logger"
	"github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

type idx int

const (
	Default idx = iota
	Session
	SessionForAdmin
	RateLimiter
)

var __redis = &_redis{
	dbs: make(map[idx]*redis.Client),
}

type _redis struct {
	dbs map[idx]*redis.Client
	mu  sync.RWMutex
}

func DB(n ...idx) *redis.Client {
	i := Default
	if len(n) > 0 {
		i = n[0]
	}
	__redis.mu.RLock()
	client, ok := __redis.dbs[i]
	if ok {
		__redis.mu.RUnlock()
		return client
	}
	__redis.mu.RUnlock()
	__redis.mu.Lock()
	defer __redis.mu.Unlock()
	if client, ok = __redis.dbs[i]; ok {
		return client
	}
	client = redis.NewClient(&redis.Options{
		Addr:         config.RedisAddress,
		Password:     config.RedisPassword,
		DB:           int(i),
		ReadTimeout:  config.RedisOptReadTimeout,
		WriteTimeout: config.RedisOptWriteTimeout,
		MaxConnAge:   config.RedisOptMaxConnAge,
		MinIdleConns: config.RedisOptMinIdleConns,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		client.Close()
		logger.Panic("failed to connect redis db%d: %v", i, err)
	}
	__redis.dbs[i] = client
	return client
}
