package kabkota

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type KabKotaServiceClient struct {
	Client pb.KabKotaServiceClient
}

func InitServiceClient(c *config.Config) pb.KabKotaServiceClient {
	cc := server.InitGRPCConn(c.TracerSvcUrl, false, "")

	return pb.NewKabKotaServiceClient(cc)
}
