package prodi

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type ProdiServiceClient struct {
	Client pb.ProdiServiceClient
}

func InitServiceClient(c *config.Config) pb.ProdiServiceClient {
	cc := server.InitGRPCConn(c.TracerSvcUrl, false, "")

	return pb.NewProdiServiceClient(cc)
}
