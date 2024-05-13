package auth

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc := server.InitGRPCConn(c.AuthSvcUrl, false, "")

	return pb.NewAuthServiceClient(cc)
}
