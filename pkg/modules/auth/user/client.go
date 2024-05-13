package user

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func InitServiceClient(c *config.Config) pb.UserServiceClient {
	cc := server.InitGRPCConn(c.AuthSvcUrl, false, "")

	return pb.NewUserServiceClient(cc)
}