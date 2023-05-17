package main

import (
	"context"
	"fmt"
	"net"

	"github.com/charmbracelet/log"
	"github.com/flexwie/svc/auth-svc/pkg/config"
	"github.com/flexwie/svc/auth-svc/pkg/services"
	"github.com/flexwie/svc/common/pkg/logger"
	"github.com/flexwie/svc/common/pkg/pb"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	fx.New(
		logger.WithFxLogger,
		logger.WithLogger,
		fx.Provide(NewConfig),
		fx.Provide(NewGrpcServer),
		fx.Invoke(func(*grpc.Server) {}),
	).Run()
}

func NewConfig(lc fx.Lifecycle, logger *log.Logger) config.Config {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed at config", err)
	}

	return c
}

func NewGrpcServer(lc fx.Lifecycle, c config.Config, logger *log.Logger) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Port))
	if err != nil {
		log.Fatalf("failed to listen at port", err)
	}

	logger.SetPrefix("grpc")
	logger.Info("svc started", "addr", lis.Addr())

	s := services.Server{
		Logger: logger.WithPrefix("grpc"),
	}

	grpcServer := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			pb.RegisterAuthServiceServer(grpcServer, &s)
			go func() {
				if err := grpcServer.Serve(lis); err != nil {
					s.Logger.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.Stop()
			return nil
		},
	})

	return grpcServer
}
