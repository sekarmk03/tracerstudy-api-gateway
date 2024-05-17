package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type GetNimByDataAtasanRequestBody struct {
	NamaAtasan  string `json:"nama_atasan"`
	EmailAtasan string `json:"email_atasan"`
	HpAtasan    string `json:"hp_atasan"`
}

func GetNimByDataAtasan(ctx *gin.Context, c pb.PKTSServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	b := GetNimByDataAtasanRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.GetNimByDataAtasan(grpcCtx, &pb.GetNimByDataAtasanRequest{
		NamaAtasan:  b.NamaAtasan,
		EmailAtasan: b.EmailAtasan,
		HpAtasan:    b.HpAtasan,
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
