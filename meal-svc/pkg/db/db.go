package db

import (
	"github.com/flexwie/svc/common/pkg/logger"
	"github.com/flexwie/svc/meal-svc/pkg/models"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormConnectionString string

func NewDb(conn GormConnectionString, logger *logger.GormLogger) *gorm.DB {
	logger.Logger.Info(conn)
	db, err := gorm.Open(sqlite.Open(string(conn)), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		logger.Logger.Fatal("could not open db", "err", err)
	}

	err = db.AutoMigrate(&models.Meal{})
	if err != nil {
		logger.Logger.Fatal(err)
	}

	return db
}

func WithDb(con string) fx.Option {
	return fx.Options(
		fx.Provide(logger.NewGormLogger),
		fx.Provide(
			fx.Annotate(
				func() GormConnectionString { return GormConnectionString(con) },
				fx.ResultTags("connection_string"),
			),
		),
		fx.Provide(
			fx.Annotate(
				NewDb,
				fx.ParamTags("connection_string", ""),
			),
		),
	)
}
