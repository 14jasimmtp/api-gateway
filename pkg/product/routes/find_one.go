package routes

import (
	"context"
	"net/http"

	"github.com/14jasimmtp/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type FindOneRequestbody struct{
	Pid int64 
}

func FindOne(c *gin.Context,p pb.ProductServiceClient){
	body:=FindOneRequestbody{}

	if err := c.BindJSON(&body); err != nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	resp,err:=p.FindOne(context.Background(),&pb.FindOneRequest{
		Id: body.Pid,
	})

	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusCreated,&resp)
}