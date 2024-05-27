package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type GetAlumniListByAtasanRequestBody struct {
	EmailAtasan string `json:"email_atasan"`
	HpAtasan    string `json:"hp_atasan"`
	NamaAtasan  string `json:"nama_atasan"`
}

func GetAlumniListByAtasan(ctx *gin.Context, c pb.UserStudyServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	b := GetAlumniListByAtasanRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.GetAlumniListByAtasan(grpcCtx, &pb.GetAlumniByAtasanRequest{
		EmailAtasan: b.EmailAtasan,
		HpAtasan:    b.HpAtasan,
		NamaAtasan:  b.NamaAtasan,
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
