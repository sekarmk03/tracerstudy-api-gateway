package post

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type PostServiceClient struct {
	Client pb.PostServiceClient
}

func InitServiceClient(c *config.Config) pb.PostServiceClient {
	cc := server.InitGRPCConn(c.PostSvcUrl, false, "")

	return pb.NewPostServiceClient(cc)
}
