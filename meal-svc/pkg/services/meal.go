package services

import (
	"context"
	"errors"

	"github.com/charmbracelet/log"
	"github.com/flexwie/svc/common/pkg/pb"
	"github.com/flexwie/svc/meal-svc/pkg/models"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedMealServiceServer
	Logger *log.Logger
	Db     *gorm.DB
}

func (s *Server) CreateMeal(ctx context.Context, req *pb.CreateMealRequest) (*pb.CreateMealResponse, error) {
	meal := models.Meal{}
	meal.FromPbRequest(req)

	result := s.Db.Create(&meal)
	if result.Error != nil {
		return &pb.CreateMealResponse{}, result.Error
	}

	var returnValue *pb.Meal
	meal.ToPbResponse(returnValue)

	return &pb.CreateMealResponse{
		Status: 1,
		Data:   returnValue,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var meal models.Meal
	var res pb.Meal

	tx := s.Db.First(&meal, req.Id)
	if tx.RowsAffected == 0 {
		return &pb.FindOneResponse{}, errors.New("could not find entry")
	}

	meal.ToPbResponse(&res)

	return &pb.FindOneResponse{
		Status: 1,
		Data:   &res,
	}, nil
}
