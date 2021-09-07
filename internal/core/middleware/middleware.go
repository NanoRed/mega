package middleware

import (
	"github.com/RedAFD/mega/internal/core/handler"
)

type Wrapper func(handle handler.Handler, options ...Option) handler.Handler

type Option struct {
	Index int
	Value interface{}
}
