package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/auth/routes"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/api/v1/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
