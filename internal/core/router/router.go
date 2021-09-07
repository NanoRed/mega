package router

import (
	"github.com/RedAFD/mega/internal/core/context"
	"github.com/RedAFD/mega/internal/core/handler"
)

type Router interface {
	Get(path string, handle handler.Handler)
	Put(path string, handle handler.Handler)
	Post(path string, handle handler.Handler)
	Group(prefix string) Router
	Handler(ctx context.Context)
	ParamInPath(name string) string
}

var NewRouter func() Router
