package pkts

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type PKTSServiceClient struct {
	Client pb.PKTSServiceClient
}

func InitServiceClient(c *config.Config) pb.PKTSServiceClient {
	cc := server.InitGRPCConn(c.TracerSvcUrl, false, "")

	return pb.NewPKTSServiceClient(cc)
}