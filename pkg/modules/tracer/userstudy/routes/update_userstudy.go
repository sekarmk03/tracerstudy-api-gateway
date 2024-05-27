package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type UpdateUserStudyRequestBody struct {
	EmailResponden                    string `json:"email_responden"`
	HpResponden                       string `json:"hp_responden"`
	NimLulusan                        string `json:"nim_lulusan"`
	LamaMengenalLulusan               uint32 `json:"lama_mengenal_lulusan"`
	Etika                             string `json:"etika"`
	KeahlianBidIlmu                   string `json:"keahlian_bid_ilmu"`
	PengembanganDiri                  string `json:"pengembangan_diri"`
	BahasaInggris                     string `json:"bahasa_inggris"`
	PenggunaanTi                      string `json:"penggunaan_ti"`
	KerjasamaTim                      string `json:"kerjasama_tim"`
	Komunikasi                        string `json:"komunikasi"`
	KesiapanTerjunMasy                string `json:"kesiapan_terjun_masy"`
	KeunggulanLulusan                 string `json:"keunggulan_lulusan"`
	KelemahanLulusan                  string `json:"kelemahan_lulusan"`
	SaranPeningkatanKompetensiLulusan string `json:"saran_peningkatan_kompetensi_lulusan"`
	SaranPerbaikanKurikulum           string `json:"saran_perbaikan_kurikulum"`
}

func UpdateUserStudy(ctx *gin.Context, c pb.UserStudyServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	b := UpdateUserStudyRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.UpdateUserStudy(grpcCtx, &pb.UserStudy{
		EmailResponden:                    b.EmailResponden,
		HpResponden:                       b.HpResponden,
		NimLulusan:                        b.NimLulusan,
		LamaMengenalLulusan:               b.LamaMengenalLulusan,
		Etika:                             b.Etika,
		KeahlianBidIlmu:                   b.KeahlianBidIlmu,
		PengembanganDiri:                  b.PengembanganDiri,
		BahasaInggris:                     b.BahasaInggris,
		PenggunaanTi:                      b.PenggunaanTi,
		KerjasamaTim:                      b.KerjasamaTim,
		Komunikasi:                        b.Komunikasi,
		KesiapanTerjunMasy:                b.KesiapanTerjunMasy,
		KeunggulanLulusan:                 b.KeunggulanLulusan,
		KelemahanLulusan:                  b.KelemahanLulusan,
		SaranPeningkatanKompetensiLulusan: b.SaranPeningkatanKompetensiLulusan,
		SaranPerbaikanKurikulum:           b.SaranPerbaikanKurikulum,
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
