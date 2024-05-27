package pkts

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type UserStudyServiceClient struct {
	Client pb.UserStudyServiceClient
}

func InitServiceClient(c *config.Config) pb.UserStudyServiceClient {
	cc := server.InitGRPCConn(c.TracerSvcUrl, false, "")

	return pb.NewUserStudyServiceClient(cc)
}
