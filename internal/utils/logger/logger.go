package logger

import (
	"github.com/RedAFD/mega/internal/utils/logger/impl"
)

type logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Panic(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}

var loggerEngine logger = impl.LoggerZap()

func Debug(format string, v ...interface{}) {
	loggerEngine.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	loggerEngine.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	loggerEngine.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	loggerEngine.Error(format, v...)
}

func Panic(format string, v ...interface{}) {
	loggerEngine.Panic(format, v...)
}

func Fatal(format string, v ...interface{}) {
	loggerEngine.Fatal(format, v...)
}
