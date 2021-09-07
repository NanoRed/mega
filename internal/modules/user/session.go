package user

import (
	"time"

	"github.com/RedAFD/mega/internal/storage/redis/impl"
)

type sessionStorage interface {
	Set(key string, value []byte, expiration time.Duration) (err error)
	Get(key string) (value []byte, err error)
}

var sessionStorageEngine sessionStorage = impl.SessionStorageRedis()

func SetSession(sessionID string, value []byte, expiration time.Duration) (err error) {
	return sessionStorageEngine.Set(sessionID, value, expiration)
}

func GetSession(sessionID string) (value []byte, err error) {
	return sessionStorageEngine.Get(sessionID)
}
