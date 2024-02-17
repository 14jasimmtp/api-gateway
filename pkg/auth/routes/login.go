package routes

import (
	"context"
	"net/http"

	"github.com/14jasimmtp/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct{
	Email string
	Password string
}

func Login(c *gin.Context, p pb.AuthServiceClient){
	body:=LoginRequestBody{}

	if err:= c.BindJSON(&body); err != nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return 
	}

	resp,err:=p.Login(context.Background(),&pb.LoginRequest{
		Email: body.Email,
		Password: body.Password,
	})

	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
		return 
	}
	c.JSON(http.StatusCreated,&resp)
}
