package product

import (
	"github.com/gin-gonic/gin"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/auth"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/config"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/api/v1/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/", svc.FindOne)
}

func (s *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, s.Client)
}

func (s *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, s.Client)
}
