package main

import (
	"log"

	"github.com/14jasimmtp/api-gateway/pkg/auth"
	"github.com/14jasimmtp/api-gateway/pkg/config"
	"github.com/14jasimmtp/api-gateway/pkg/order"
	"github.com/14jasimmtp/api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	
	authSvc := *auth.RegisterRoutes(r, &c)
	order.OrderRoutes(r, &c, &authSvc)
	product.ProductRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
