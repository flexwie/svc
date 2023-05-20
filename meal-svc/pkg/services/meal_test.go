package services_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/flexwie/svc/common/pkg/pb"
	"github.com/flexwie/svc/meal-svc/pkg/models"
	"github.com/flexwie/svc/meal-svc/pkg/services"
	"github.com/stretchr/testify/assert"
	"gopkg.in/check.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test(t *testing.T) { check.TestingT(t) }

type TestSuite struct {
	s   *services.Server
	log bytes.Buffer
	db  *gorm.DB
}

var _ = check.Suite(&TestSuite{})

func (s *TestSuite) SetUpSuite(c *check.C) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("can not create db")
	}

	s.db = db

	s.s = &services.Server{
		Logger: log.New(&s.log),
		Db:     db,
	}
}

func (s *TestSuite) SetUpTest(c *check.C) {
	s.db.AutoMigrate(models.Meal{})
}

func (s *TestSuite) TearDownTest(c *check.C) {
	s.db.Delete(&models.Meal{})
}

func (s *TestSuite) TestGetFailed(c *check.C) {
	_, err := s.s.FindOne(context.TODO(), &pb.FindOneRequest{Id: "100"})

	c.Assert(err, check.NotNil)
}

func (s *TestSuite) TestGet(c *check.C) {
	meal := &models.Meal{Meal: pb.Meal{Name: "test"}}
	s.db.Create(meal)
	res, err := s.s.FindOne(context.TODO(), &pb.FindOneRequest{Id: fmt.Sprint(meal.ID)})

	assert.NoError(c, err)
	assert.Equal(c, res.Data.Name, "test")
}

// func (s *TestSuite) TestCreate(c *check.C) {
// 	res, err := s.s.CreateMeal(context.TODO(), &pb.CreateMealRequest{Name: "Create"})
//
// 	assert.NoError(c, err)
// 	assert.Equal(c, res.Data.Id, "1")
// }
