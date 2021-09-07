package config

import "time"

var (
	AppDebug   bool   = true
	AppLogDir  string = ""
	AppLang    string = "zh-CN"
	AppListen  string = "0.0.0.0:80"
	AppWebHost string = "www.yoursite.com"

	PostgresDSN                      string        = "postgres://username:123456@127.0.0.1:5432/dbname"
	PostgresOptReadTimeout           time.Duration = time.Second * 3
	PostgresOptWriteTimeout          time.Duration = time.Second * 3
	PostgresOptMaxConnAge            time.Duration = time.Minute * 2
	PostgresOptMinIdleConns          int           = 3
	PostgresOptMaxRetries            int           = 2
	PostgresOptRetryStatementTimeout bool          = true

	RedisAddress                   string        = "127.0.0.1:6379"
	RedisPassword                  string        = ""
	RedisOptReadTimeout            time.Duration = time.Second * 3
	RedisOptWriteTimeout           time.Duration = time.Second * 3
	RedisOptMaxConnAge             time.Duration = time.Minute * 2
	RedisOptMinIdleConns           int           = 3
	RedisDefaultExpiration         time.Duration = time.Hour
	RedisExpirationSecondSaltRange int           = 3600

	SessionExpiration         time.Duration = time.Hour * 24 * 90
	SessionForAdminExpiration time.Duration = time.Hour

	GocacheDefaultExpiration time.Duration = time.Minute * 5
	GocacheCleanupInterval   time.Duration = time.Minute * 10

	EmailDefaultSenderUsername string = ""
	EmailDefaultSenderPassword string = ""
	EmailDefaultSenderHost     string = ""
	EmailDefaultSenderPort     string = ""
	EmailConfirmLinkFormat     string = ""

	UserPasswordCost int = 10

	RateLimitPeriod    time.Duration = time.Hour
	RateLimitFrequency int           = 1000
)
