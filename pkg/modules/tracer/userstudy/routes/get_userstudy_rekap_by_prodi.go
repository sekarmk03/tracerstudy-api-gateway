package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func GetUserStudyRekapByProdi(ctx *gin.Context, c pb.UserStudyServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	kode := ctx.Param("kode")

	res, err := c.GetUserStudyRekapByProdi(grpcCtx, &pb.GetUserStudyRekapByProdiRequest{
		KodeProdi: kode,
	})

	if err != nil {
		errResp := utils.GetGrpcError(err)
		ctx.AbortWithStatusJSON(
			errResp.Code,
			errResp,
		)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
