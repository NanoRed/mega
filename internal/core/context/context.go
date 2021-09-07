package context

import (
	c "context"
	"time"
)

type Context interface {
	c.Context
	GetReqIP() string
	GetReqPath() []byte
	GetReqBody() []byte
	GetReqCookie(key string) []byte
	SetRespCode(code int, content ...interface{})
	SetRespRedirect(url string, code int)
	SetRespHeader(key, value string)
	SetRespCookie(key, value string, expire time.Time)
	SetRespContentType(contentType string)
	GetParam(key string) interface{}
	SetParam(key string, value interface{})
}
