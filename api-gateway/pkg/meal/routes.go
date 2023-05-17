package meal

import (
	"github.com/charmbracelet/log"
	"github.com/flexwie/svc/gateway/pkg/auth"
	"github.com/flexwie/svc/gateway/pkg/config"
	"github.com/flexwie/svc/gateway/pkg/meal/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, auth *auth.AuthMiddlewareConfig, logger *log.Logger) {
	svc := &ServiceClient{
		Client: InitServiceClient(c, logger),
		Logger: logger.WithPrefix("meal"),
	}

	routes := r.Group("/meal")
	routes.Use(auth.AuthRequired)
	routes.POST("/", svc.CreateMeal)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) CreateMeal(ctx *gin.Context) {
	svc.Logger.Debug("creating meal")
	routes.CreateMeal(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	svc.Logger.Debug("finding meal")
	routes.FindOne(ctx, svc.Client, svc.Logger)
}
