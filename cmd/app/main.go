package main

import (
	"log"
	"time"

	"github.com/RedAFD/mega/internal/config"
	"github.com/RedAFD/mega/internal/core"
	userApi "github.com/RedAFD/mega/internal/modules/user/api"
	"github.com/RedAFD/mega/internal/utils/rate"
	"github.com/RedAFD/mega/third_party/swagger"
)

// @Title web APIs
// @Version 1.0
// @Contact.name RedAFD
// @Contact.email radixholms@gmail.com
// @Host www.yoursite.com
// @Host 127.0.0.1
// @BasePath /api/v1

func main() {

	// use UTC as timezone
	time.Local = time.UTC

	// swagger
	if config.AppDebug {
		core.Route("GET", "/swagger/"+core.PInP("*"), swagger.Handler)
	}

	// api routes
	api := core.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.Post("/login", userApi.Login)
			}
		}
	}

	// core.Before() should be called after core.Route()
	core.Before(rate.Limit,
		rate.WithPeriod(config.RateLimitPeriod),
		rate.WithLimit(config.RateLimitFrequency),
	)

	log.Fatal(core.Run())
}
