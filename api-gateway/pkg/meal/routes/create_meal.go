package routes

import (
	"context"
	"net/http"

	"github.com/flexwie/svc/common/pkg/pb"
	"github.com/gin-gonic/gin"
)

type CreateMealRequestBody struct {
	Title string `json:"title"`
}

func CreateMeal(ctx *gin.Context, c pb.MealServiceClient) {
	body := CreateMealRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//userId, _ := ctx.Get("userId")

	res, err := c.CreateMeal(context.Background(), &pb.CreateMealRequest{
		Name: body.Title,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
