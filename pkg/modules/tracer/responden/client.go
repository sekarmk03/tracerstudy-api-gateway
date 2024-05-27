package responden

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type RespondenServiceClient struct {
	Client pb.RespondenServiceClient
}

func InitServiceClient(c *config.Config) pb.RespondenServiceClient {
	cc := server.InitGRPCConn(c.TracerSvcUrl, false, "")

	return pb.NewRespondenServiceClient(cc)
}
