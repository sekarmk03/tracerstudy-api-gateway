package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type CreateUserStudyRequestBody struct {
	NamaResponden  string `json:"nama_responden"`
	EmailResponden string `json:"email_responden"`
	HpResponden    string `json:"hp_responden"`
	NamaInstansi   string `json:"nama_instansi"`
	Jabatan        string `json:"jabatan"`
	AlamatInstansi string `json:"alamat_instansi"`
	NimLulusan     string `json:"nim_lulusan"`
	NamaLulusan    string `json:"nama_lulusan"`
	ProdiLulusan   string `json:"prodi_lulusan"`
	TahunLulusan   string `json:"tahun_lulusan"`
}

func CreateUserStudy(ctx *gin.Context, c pb.UserStudyServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	b := CreateUserStudyRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.CreateUserStudy(grpcCtx, &pb.UserStudy{
		NamaResponden:  b.NamaResponden,
		EmailResponden: b.EmailResponden,
		HpResponden:    b.HpResponden,
		NamaInstansi:   b.NamaInstansi,
		Jabatan:        b.Jabatan,
		AlamatInstansi: b.AlamatInstansi,
		NimLulusan:     b.NimLulusan,
		NamaLulusan:    b.NamaLulusan,
		ProdiLulusan:   b.ProdiLulusan,
		TahunLulusan:   b.TahunLulusan,
	})

	if err != nil {
		errResp := utils.GetGrpcError(err)
		ctx.AbortWithStatusJSON(
			errResp.Code,
			errResp,
		)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
