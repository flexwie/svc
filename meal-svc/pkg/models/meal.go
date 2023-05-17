package models

import (
	"fmt"

	"github.com/flexwie/svc/common/pkg/pb"
	"gorm.io/gorm"
)

type Meal struct {
	gorm.Model
	pb.Meal
}

func (m *Meal) FromPbRequest(req *pb.CreateMealRequest) {
	m.Name = req.Name
}

func (m *Meal) ToPbResponse(req *pb.Meal) {
	req.Id = fmt.Sprint(m.ID)
	req.Name = m.Name
}
