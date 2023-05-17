package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/flexwie/svc/common/pkg/pb"
	"github.com/gin-gonic/gin"
)

func FindOne(ctx *gin.Context, c pb.MealServiceClient, logger *log.Logger) {
	id := ctx.Param("id")

	if id == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("id must be set in params"))
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: id,
	})

	if err != nil {
		logger.Error("could not find meal", "err", err, "id", id)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
