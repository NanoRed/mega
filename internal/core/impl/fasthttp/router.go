package impl

import (
	"bytes"

	"github.com/RedAFD/mega/internal/core/context"
	"github.com/RedAFD/mega/internal/core/handler"
	"github.com/RedAFD/mega/internal/core/router"
	frouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func init() {
	router.NewRouter = func() router.Router {
		return &_coreRouterFasthttp{frouter.New(), ""}
	}
}

type _coreRouterFasthttp struct {
	*frouter.Router
	prefix string
}

func (c *_coreRouterFasthttp) Get(path string, handle handler.Handler) {
	c.GET(c.prefix+path, fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		handle(newCoreContextFasthttp(ctx))
	}))
}

func (c *_coreRouterFasthttp) Put(path string, handle handler.Handler) {
	c.PUT(c.prefix+path, fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		handle(newCoreContextFasthttp(ctx))
	}))
}

func (c *_coreRouterFasthttp) Post(path string, handle handler.Handler) {
	c.POST(c.prefix+path, fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		handle(newCoreContextFasthttp(ctx))
	}))
}

func (c *_coreRouterFasthttp) Group(prefix string) router.Router {
	return &_coreRouterFasthttp{c.Router, c.prefix + prefix}
}

func (c *_coreRouterFasthttp) Handler(ctx context.Context) {
	c.Router.Handler(ctx.(*_coreContextFasthttp).RequestCtx)
}

func (c *_coreRouterFasthttp) ParamInPath(name string) string {
	if name == "*" {
		return "{filepath:*}"
	}
	buf := &bytes.Buffer{}
	buf.WriteByte('{')
	buf.WriteString(name)
	buf.WriteByte('}')
	return buf.String()
}
