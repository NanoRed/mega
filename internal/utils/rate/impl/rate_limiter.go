package impl

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/RedAFD/mega/internal/core/context"
	"github.com/RedAFD/mega/internal/core/handler"
	"github.com/RedAFD/mega/internal/storage/redis"
	"github.com/RedAFD/mega/internal/utils/i18n"
	"github.com/RedAFD/mega/internal/utils/logger"
	"github.com/ulule/limiter/v3"
	limiterRedisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

var __rateLimiter = &_rateLimiter{}

type _rateLimiter struct {
	store *limiter.Store
	once  sync.Once
}

func (r *_rateLimiter) Limit(
	handle handler.Handler,
	period time.Duration,
	limit int,
	keyGetter func(ctx context.Context) (key string, exclude bool),
) handler.Handler {

	limiter := limiter.New(*r.store, limiter.Rate{
		Period: period,
		Limit:  int64(limit),
	})

	return handler.Handler(func(ctx context.Context) {

		key, exclude := keyGetter(ctx)

		if exclude {
			handle(ctx)
			return
		}

		context, err := limiter.Get(ctx, key)
		if err != nil {
			logger.Error("rate limit err: %v", err)
			handle(ctx)
			return
		}

		ctx.SetRespHeader("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
		ctx.SetRespHeader("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		ctx.SetRespHeader("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

		if context.Reached {
			ctx.SetRespCode(http.StatusTooManyRequests, i18n.Sprintf("服务器繁忙，请稍后再试"))
			return
		}

		handle(ctx)
	})
}

func RateLimiter() *_rateLimiter {
	__rateLimiter.once.Do(func() {
		store, err := limiterRedisStore.NewStore(redis.DB(redis.RateLimiter))
		if err != nil {
			logger.Panic("failed to new store: %v", err)
		}
		__rateLimiter.store = &store
	})
	return __rateLimiter
}
