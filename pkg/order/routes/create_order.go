package routes

import (
	"context"
	"net/http"

	"github.com/14jasimmtp/api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct{
	ProductId int64
	Quantity int64
	OrderId int64
}

func CreateOrder(c *gin.Context,p pb.OrderServiceClient){
	body:=CreateOrderRequest{}
	if err := c.BindJSON(&body);err != nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	resp,err:=p.CreateOrder(context.Background(),&pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity: body.Quantity,
		OrderId: body.OrderId,
	})

	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusOK,&resp)
}