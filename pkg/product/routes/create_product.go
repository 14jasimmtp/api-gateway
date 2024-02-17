package routes

import (
	"context"
	"net/http"

	"github.com/14jasimmtp/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct{
	Name string
	Sku string
	Stock int
	Price float64
}

func CreateProduct(c *gin.Context,p pb.ProductServiceClient){
	body:=CreateProductRequest{}

	if err := c.BindJSON(&body);err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	resp,err:=p.CreateProduct(context.Background(),&pb.CreateProductRequest{
		Name: body.Name,
		Sku: body.Sku,
		Stock: int64(body.Stock),
		Price: int64(body.Price),
	})

	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusCreated,&resp)
}