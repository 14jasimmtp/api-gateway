package routes

import (
	"context"
	"net/http"

	"github.com/14jasimmtp/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct{
	Email string
	Password string
}

func Register(c *gin.Context, p pb.AuthServiceClient){
	body:=RegisterRequestBody{}

	if err:= c.BindJSON(&body); err != nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return 
	}

	resp,err:=p.Register(context.Background(),&pb.RegisterRequest{
		Email: body.Email,
		Password: body.Password,
	})

	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
		return 
	}
	c.JSON(int(resp.Status),&resp)
}
