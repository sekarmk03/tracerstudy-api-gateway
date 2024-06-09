package comment

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/server"
)

type CommentServiceClient struct {
	Client pb.CommentServiceClient
}

func InitServiceClient(c *config.Config) pb.CommentServiceClient {
	cc := server.InitGRPCConn(c.PostSvcUrl, false, "")

	return pb.NewCommentServiceClient(cc)
}
