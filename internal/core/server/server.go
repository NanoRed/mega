package server

import (
	"github.com/RedAFD/mega/internal/core/handler"
)

type Server interface {
	ListenAndServe(addr string, handler handler.Handler) error
}

var NewServer func() Server
