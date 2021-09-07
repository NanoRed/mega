package rate

import (
	"time"

	"github.com/RedAFD/mega/internal/core/context"
	"github.com/RedAFD/mega/internal/core/handler"
	"github.com/RedAFD/mega/internal/core/middleware"
	"github.com/RedAFD/mega/internal/utils/rate/impl"
)

type rate interface {
	Limit(
		handle handler.Handler,
		period time.Duration,
		limit int,
		keyGetter func(ctx context.Context) (
			key string, // this value should be used as a key for frequency count
			exclude bool, // when this value is true, the limiter will be bypassed
		),
	) handler.Handler
}

var rateEngine rate = impl.RateLimiter()

const (
	_ = iota
	optionPeriod
	optionLimit
	optionKeyGetter
)

var defaultKeyGetter = func(ctx context.Context) (key string, exclude bool) {
	key = ctx.GetReqIP()
	return
}

func Limit(handle handler.Handler, options ...middleware.Option) handler.Handler {

	var period time.Duration
	var limit int
	var keyGetter = defaultKeyGetter

	require := 2
	for _, option := range options {
		switch option.Index {
		case optionPeriod:
			period = option.Value.(time.Duration)
			require--
		case optionLimit:
			limit = option.Value.(int)
			require--
		case optionKeyGetter:
			keyGetter = option.Value.(func(ctx context.Context) (key string, exclude bool))
		}
	}
	if require > 0 {
		return handle
	}

	return rateEngine.Limit(handle, period, limit, keyGetter)
}

func WithPeriod(period time.Duration) middleware.Option {
	return middleware.Option{
		Index: optionPeriod,
		Value: period,
	}
}

func WithLimit(limit int) middleware.Option {
	return middleware.Option{
		Index: optionLimit,
		Value: limit,
	}
}

func WithKeyGetter(keyGetter func(ctx context.Context) (key string, exclude bool)) middleware.Option {
	return middleware.Option{
		Index: optionKeyGetter,
		Value: keyGetter,
	}
}
