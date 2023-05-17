package main

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/flexwie/svc/common/pkg/logger"
	"github.com/flexwie/svc/gateway/pkg/auth"
	"github.com/flexwie/svc/gateway/pkg/config"
	"github.com/flexwie/svc/gateway/pkg/meal"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		logger.WithFxLogger,
		logger.WithLogger,
		auth.Module,
		fx.Provide(NewConfig),
		fx.Provide(NewRouter),
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
}

func NewConfig(logger *log.Logger) config.Config {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatal("failed at config", err)
	}

	return c
}

func NewRouter(lc fx.Lifecycle, c config.Config, a *auth.AuthMiddlewareConfig, clog *log.Logger) *gin.Engine {
	clog.WithPrefix("gin")

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(logger.NewGinLogger(clog))
	r.Use(gin.Recovery())

	meal.RegisterRoutes(r, &c, a, clog)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go r.Run(c.Port)

			return nil
		},
	})

	return r
}
