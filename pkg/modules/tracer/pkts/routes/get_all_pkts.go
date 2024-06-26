package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func GetAllPKTS(ctx *gin.Context, c pb.PKTSServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	limitStr := ctx.DefaultQuery("limit", "10")
	pageStr := ctx.DefaultQuery("page", "1")

	limit := utils.StrParamToInt(limitStr, 10)
	page := utils.StrParamToInt(pageStr, 1)

	res, err := c.GetAllPKTS(grpcCtx, &pb.GetAllPKTSRequest{
		Pagination: &pb.PaginationRequest{
			Limit: uint32(limit),
			Page:  uint32(page),
		},
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
