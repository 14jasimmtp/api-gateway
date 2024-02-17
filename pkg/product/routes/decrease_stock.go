package routes

import (
	"context"
	"net/http"

	"github.com/14jasimmtp/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type DecreaseStockRequestbody struct{
	Pid int64
	OdId int64
}

func DecreaseStock(c *gin.Context,p pb.ProductServiceClient){
	body:=DecreaseStockRequestbody{}

	if err := c.BindJSON(&body); err!= nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	resp,err:=p.DecreaseStock(context.Background(),&pb.DecreaseStockRequest{
		Id: body.Pid,
		OrderId: body.OdId,
	})

	if err != nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusCreated,&resp)
}