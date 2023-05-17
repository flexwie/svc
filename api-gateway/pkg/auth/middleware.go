package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/flexwie/svc/common/pkg/pb"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

var Module = fx.Options(
	fx.Provide(NewServiceClient),
	fx.Provide(NewMiddleware),
)

func NewMiddleware(svc *ServiceClient) *AuthMiddlewareConfig {
	return &AuthMiddlewareConfig{svc: svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	ctx.Next()
	return

	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Next()
}
