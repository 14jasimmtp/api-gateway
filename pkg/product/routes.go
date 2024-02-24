package product

import (
	"github.com/14jasimmtp/api-gateway/pkg/auth"
	"github.com/14jasimmtp/api-gateway/pkg/config"
	"github.com/14jasimmtp/api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/add-product", svc.CreateProduct)
	routes.GET("/find", svc.FindOne)
}

func (svc *ServiceClient) CreateProduct(c *gin.Context) {
	routes.CreateProduct(c, svc.Client)
}

func (svc *ServiceClient) DecreaseStock(c *gin.Context) {
	routes.DecreaseStock(c, svc.Client)
}

func (svc *ServiceClient) FindOne(c *gin.Context) {
	routes.FindOne(c, svc.Client)
}
