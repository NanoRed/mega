package core

import (
	"net/http"

	"github.com/RedAFD/mega/internal/config"
	"github.com/RedAFD/mega/internal/core/handler"
	_ "github.com/RedAFD/mega/internal/core/impl/fasthttp"
	"github.com/RedAFD/mega/internal/core/middleware"
	"github.com/RedAFD/mega/internal/core/router"
	"github.com/RedAFD/mega/internal/core/server"
)

var DefaultRouter = router.NewRouter()

func Group(prefix string) router.Router {
	return DefaultRouter.Group(prefix)
}

func Route(method, path string, handle handler.Handler) {
	switch method {
	case http.MethodGet:
		DefaultRouter.Get(path, handle)
	case http.MethodPut:
		DefaultRouter.Put(path, handle)
	case http.MethodPost:
		DefaultRouter.Post(path, handle)
	}
}

func PInP(name string) string {
	return DefaultRouter.ParamInPath(name)
}

func Before(wrapper middleware.Wrapper, options ...middleware.Option) {
	wrapper(DefaultRouter.Handler, options...)
}

var DefaultServer = server.NewServer()

func Run() error {
	return DefaultServer.ListenAndServe(config.AppListen, DefaultRouter.Handler)
}
