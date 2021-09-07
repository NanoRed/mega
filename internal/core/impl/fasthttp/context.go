package impl

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type _coreContextFasthttp struct {
	*fasthttp.RequestCtx
}

func newCoreContextFasthttp(ctx *fasthttp.RequestCtx) *_coreContextFasthttp {
	return &_coreContextFasthttp{ctx}
}

func (c *_coreContextFasthttp) GetReqIP() string {
	return c.RemoteIP().String()
}

func (c *_coreContextFasthttp) GetReqPath() []byte {
	return c.Path()
}

func (c *_coreContextFasthttp) GetReqBody() []byte {
	return c.Request.Body()
}

func (c *_coreContextFasthttp) GetReqCookie(key string) []byte {
	return c.Request.Header.Cookie(key)
}

func (c *_coreContextFasthttp) SetRespCode(code int, content ...interface{}) {
	if code/100 > 3 {
		c.Response.Reset()
	}
	c.SetStatusCode(code)
	if len(content) > 0 {
		switch v := content[0].(type) {
		case string:
			c.SetBodyString(v)
			c.SetContentType("text/plain; charset=utf-8")
		case []byte:
			c.SetBody(v)
			c.SetContentType("text/plain; charset=utf-8")
		case func(params ...interface{}) []byte:
			c.SetBody(v(content[1:]...))
		default:
			jsonb, _ := jsoniter.Marshal(v)
			c.SetBody(jsonb)
			c.SetContentType("application/json; charset=utf-8")
		}
	}
}

func (c *_coreContextFasthttp) SetRespRedirect(url string, code int) {
	c.Redirect(url, code)
}

func (c *_coreContextFasthttp) SetRespHeader(key, value string) {
	c.Response.Header.Set(key, value)
}

func (c *_coreContextFasthttp) SetRespCookie(key, value string, expire time.Time) {
	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)
	cookie.SetKey(key)
	cookie.SetValue(value)
	cookie.SetExpire(expire)
	c.Response.Header.SetCookie(cookie)
}

func (c *_coreContextFasthttp) SetRespContentType(contentType string) {
	c.SetContentType(contentType)
}

func (c *_coreContextFasthttp) GetParam(key string) interface{} {
	return c.UserValue(key)
}

func (c *_coreContextFasthttp) SetParam(key string, value interface{}) {
	c.SetUserValue(key, value)
}
