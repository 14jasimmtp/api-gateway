package order

import (
	"github.com/14jasimmtp/api-gateway/pkg/auth"
	"github.com/14jasimmtp/api-gateway/pkg/order/routes"
	"github.com/14jasimmtp/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("",svc.CreateOrder)
}

	func( svc *ServiceClient) CreateOrder(c *gin.Context){
		routes.CreateOrder(c,svc.Client)
	}

