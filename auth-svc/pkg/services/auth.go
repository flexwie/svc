package services

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/flexwie/svc/common/pkg/pb"
)

type Server struct {
	Logger *log.Logger
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	s.Logger.Info("validating token")
	// jwksUrl := "https://login.microsoftonline.com/common/discovery/v2.0/keys"

	// keySet, err := jwk.Fetch(ctx, jwksUrl)

	// if err != nil {
	// 	return &pb.ValidateResponse{Error: err.Error()}, err
	// }

	// _, err = jwt.Parse([]byte(req.Token), jwt.WithKeySet(keySet))

	// // check stuff on the parsed token

	// if err != nil {
	// 	return nil, err
	// }

	return &pb.ValidateResponse{}, nil
}
