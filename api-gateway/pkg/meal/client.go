package meal

import (
	"github.com/charmbracelet/log"
	"github.com/flexwie/svc/common/pkg/pb"
	"github.com/flexwie/svc/gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.MealServiceClient
	Logger *log.Logger
}

func InitServiceClient(c *config.Config, logger *log.Logger) pb.MealServiceClient {
	cc, err := grpc.Dial(c.MealSvcUrl, grpc.WithInsecure())

	if err != nil {
		logger.Infof("could not connect: %v", err)
	}

	return pb.NewMealServiceClient(cc)
}
