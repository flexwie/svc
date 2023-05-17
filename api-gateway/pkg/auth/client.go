package auth

import (
	"fmt"

	"github.com/flexwie/svc/common/pkg/pb"
	"github.com/flexwie/svc/gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func NewServiceClient(c config.Config) *ServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("could not connect: ", err)
	}

	client := pb.NewAuthServiceClient(cc)

	return &ServiceClient{
		Client: client,
	}
}
