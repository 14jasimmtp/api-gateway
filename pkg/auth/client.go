package auth

import (
	"fmt"

	"github.com/14jasimmtp/api-gateway/pkg/auth/pb"
	"github.com/14jasimmtp/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("could not connect : ", err)
	}

	return pb.NewAuthServiceClient(cc)
}
