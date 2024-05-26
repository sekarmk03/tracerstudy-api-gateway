package provinsi

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type ProvinsiServiceClient struct {
	Client pb.ProvinsiServiceClient
}

func InitServiceClient(c *config.Config) pb.ProvinsiServiceClient {
	cc := server.InitGRPCConn(c.TracerSvcUrl, false, "")

	return pb.NewProvinsiServiceClient(cc)
}
