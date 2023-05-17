package main

import (
	"context"
	"fmt"
	"net"

	"github.com/charmbracelet/log"
	cgrpc "github.com/flexwie/svc/common/pkg/grpc"
	"github.com/flexwie/svc/common/pkg/logger"
	"github.com/flexwie/svc/common/pkg/pb"
	"github.com/flexwie/svc/meal-svc/pkg/config"
	"github.com/flexwie/svc/meal-svc/pkg/db"
	"github.com/flexwie/svc/meal-svc/pkg/services"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		logger.WithFxLogger,
		logger.WithLogger,
		cgrpc.WithUnaryLoggingInterceptorFactory,
		db.WithDb("test.meal.db"),
		fx.Provide(NewConfig),
		fx.Provide(fx.Annotate(
			NewGrpcServer,
			fx.ParamTags("", "", "", "unary_logging", ""), // empty tags tell the di framework that those parameters are not annotated
		)),
		fx.Invoke(func(*grpc.Server) {}),
	).Run()
}

// reads the config file from the config package
func NewConfig(logger *log.Logger) config.Config {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config", "err", err)
	}

	return c
}

// creates a new instance of our grpc server
func NewGrpcServer(lc fx.Lifecycle, logger *log.Logger, c config.Config, uli grpc.UnaryServerInterceptor, db *gorm.DB) *grpc.Server {
	// get port to listen on
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Port))
	if err != nil {
		log.Fatal("failed to listen", "err", err, "port", c.Port)
	}

	// initiate the service
	s := services.Server{
		Logger: logger.WithPrefix("grpc"),
		Db:     db,
	}

	// create the grpc server with logging interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(uli),
	)

	// add server to lifecycle hooks
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			pb.RegisterMealServiceServer(grpcServer, &s)
			go func() {
				if err := grpcServer.Serve(lis); err != nil {
					s.Logger.Fatal(err)
				}
			}()

			s.Logger.Info("svc started", "addr", lis.Addr())

			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.Stop()
			return nil
		},
	})

	return grpcServer
}
