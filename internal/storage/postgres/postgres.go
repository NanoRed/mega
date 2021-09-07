package postgres

import (
	"context"
	"sync"

	"github.com/RedAFD/mega/internal/config"
	"github.com/RedAFD/mega/internal/utils/logger"
	"github.com/go-pg/pg/v10"
)

var __postgres = &_postgres{}

type _postgres struct {
	db   *pg.DB
	once sync.Once
}

func DB() *pg.DB {
	__postgres.once.Do(func() {
		opt, err := pg.ParseURL(config.PostgresDSN)
		if err != nil {
			logger.Panic("failed to parse postgres dsn: %v", err)
		}
		opt.ReadTimeout = config.PostgresOptReadTimeout
		opt.WriteTimeout = config.PostgresOptWriteTimeout
		opt.MaxConnAge = config.PostgresOptMaxConnAge
		opt.MinIdleConns = config.PostgresOptMinIdleConns
		opt.MaxRetries = config.PostgresOptMaxRetries
		opt.RetryStatementTimeout = config.PostgresOptRetryStatementTimeout
		__postgres.db = pg.Connect(opt)
		if err := __postgres.db.Ping(context.Background()); err != nil {
			__postgres.db.Close()
			logger.Panic("failed to connect postgres: %v", err)
		}
		if config.AppDebug {
			__postgres.db.AddQueryHook(debugHook{})
		}
	})
	return __postgres.db
}

type debugHook struct {
}

func (d debugHook) BeforeQuery(ctx context.Context, evt *pg.QueryEvent) (context.Context, error) {
	q, err := evt.FormattedQuery()
	if err != nil {
		return nil, err
	}
	if evt.Err != nil {
		logger.Debug("Postgresql executing SQL:[%s] err:[%v]", q, evt.Err)
	} else {
		// If you need this debug log, just uncomment
		logger.Debug("Postgresql executing SQL:[%s] err:[<nil>]", q)
	}
	return ctx, nil
}

func (d debugHook) AfterQuery(context.Context, *pg.QueryEvent) error {
	return nil
}
