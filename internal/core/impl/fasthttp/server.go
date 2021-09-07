package impl

import (
	"github.com/RedAFD/mega/internal/core/handler"
	"github.com/RedAFD/mega/internal/core/server"
	"github.com/valyala/fasthttp"
)

func init() {
	server.NewServer = func() server.Server {
		return &_coreServerFasthttp{&fasthttp.Server{Name: "RedAFD"}}
	}
}

type _coreServerFasthttp struct {
	*fasthttp.Server
}

func (c *_coreServerFasthttp) ListenAndServe(addr string, handler handler.Handler) error {
	c.Handler = fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		handler(newCoreContextFasthttp(ctx))
	})
	return c.Server.ListenAndServe(addr)
}
