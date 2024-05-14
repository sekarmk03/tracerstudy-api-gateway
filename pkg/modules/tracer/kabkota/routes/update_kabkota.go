package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type UpdateKabKotaRequestBody struct {
	IdWilayah      string `json:"id_wilayah"`
	Nama           string `json:"nama"`
	IdIndukWilayah string `json:"id_induk_wilayah"`
}

func UpdateKabKota(ctx *gin.Context, c pb.KabKotaServiceClient) {
	idWil := ctx.Param("id")

	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	b := UpdateKabKotaRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.UpdateKabKota(grpcCtx, &pb.KabKota{
		IdWil:      idWil,
		Nama:       b.Nama,
		IdIndukWil: b.IdIndukWilayah,
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
