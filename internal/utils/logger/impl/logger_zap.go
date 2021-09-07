package impl

import (
	"fmt"
	"os"
	"sync"

	"github.com/RedAFD/mega/internal/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var __loggerZap = &_loggerZap{}

type _loggerZap struct {
	logger *zap.SugaredLogger
	once   sync.Once
}

func (l *_loggerZap) Debug(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l *_loggerZap) Info(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *_loggerZap) Warn(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

func (l *_loggerZap) Error(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

func (l *_loggerZap) Panic(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}

func (l *_loggerZap) Fatal(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

func LoggerZap() *_loggerZap {
	__loggerZap.once.Do(func() {
		var outlog zapcore.WriteSyncer
		var errlog zapcore.WriteSyncer
		if config.AppLogDir == "" {
			outlog = zapcore.AddSync(os.Stdout)
			errlog = zapcore.AddSync(os.Stderr)
		} else {
			outlog = zapcore.AddSync(&lumberjack.Logger{
				Filename: fmt.Sprintf("%s/out.log", config.AppLogDir),
				MaxSize:  300,
				MaxAge:   7,
			})
			errlog = zapcore.AddSync(&lumberjack.Logger{
				Filename: fmt.Sprintf("%s/err.log", config.AppLogDir),
				MaxSize:  300,
				MaxAge:   7,
			})
		}
		var zl *zap.Logger
		if config.AppDebug {
			core := zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
				zapcore.AddSync(outlog),
				zap.NewAtomicLevelAt(zap.DebugLevel),
			)
			zl = zap.New(core, zap.ErrorOutput(zapcore.AddSync(errlog)), zap.AddCallerSkip(2), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		} else {
			core := zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.AddSync(outlog),
				zap.NewAtomicLevelAt(zap.InfoLevel),
			)
			zl = zap.New(core, zap.ErrorOutput(zapcore.AddSync(errlog)), zap.AddCallerSkip(2))
		}
		__loggerZap.logger = zl.Sugar()
	})
	return __loggerZap
}
