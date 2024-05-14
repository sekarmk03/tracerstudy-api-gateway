package routes

import (
	"context"
	"tracerstudy-api-gateway/pkg/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetAllKabkota(ctx *gin.Context, c pb.KabKotaServiceClient) {
	res, err := c.GetAllKabKota(context.Background(), &emptypb.Empty{})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}