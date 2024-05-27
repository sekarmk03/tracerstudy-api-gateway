package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type UpdateProdiRequestBody struct {
	Nama            string `json:"nama"`
	KodeIntegrasi   string `json:"kode_integrasi"`
	KodeDikti       string `json:"kode_dikti"`
	Jenjang         string `json:"jenjang"`
	KodeFakultas    string `json:"kode_fakultas"`
	NamaFakultas    string `json:"nama_fakultas"`
	AkronimFakultas string `json:"akronim_fakultas"`
}

func UpdateProdi(ctx *gin.Context, c pb.ProdiServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	kode := ctx.Param("kode")

	b := UpdateProdiRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.UpdateProdi(grpcCtx, &pb.Prodi{
		Kode:            kode,
		Nama:            b.Nama,
		KodeIntegrasi:   b.KodeIntegrasi,
		KodeDikti:       b.KodeDikti,
		Jenjang:         b.Jenjang,
		KodeFakultas:    b.KodeFakultas,
		NamaFakultas:    b.NamaFakultas,
		AkronimFakultas: b.AkronimFakultas,
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
