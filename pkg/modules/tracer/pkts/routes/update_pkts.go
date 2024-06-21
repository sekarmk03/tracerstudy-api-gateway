package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type UpdatePKTSRequestBody struct {
	F8                  uint32 `json:"f8"`
	F504                uint32 `json:"f504"`
	F502                uint32 `json:"f502"`
	F506                uint32 `json:"f506"`
	F505                uint64 `json:"f505"`
	F5a1                string `json:"f5a1"`
	F5a2                string `json:"f5a2"`
	F1101               uint32 `json:"f1101"`
	F1102               string `json:"f1102"`
	F5b                 string `json:"f5b"`
	F5c                 uint32 `json:"f5c"`
	F5d                 uint32 `json:"f5d"`
	F18a                uint32 `json:"f18a"`
	F18b                string `json:"f18b"`
	F18c                string `json:"f18c"`
	F18d                string `json:"f18d"`
	F1201               uint32 `json:"f1201"`
	F1202               string `json:"f1202"`
	F14                 uint32 `json:"f14"`
	F15                 uint32 `json:"f15"`
	F1761               uint32 `json:"f1761"`
	F1762               uint32 `json:"f1762"`
	F1763               uint32 `json:"f1763"`
	F1764               uint32 `json:"f1764"`
	F1765               uint32 `json:"f1765"`
	F1766               uint32 `json:"f1766"`
	F1767               uint32 `json:"f1767"`
	F1768               uint32 `json:"f1768"`
	F1769               uint32 `json:"f1769"`
	F1770               uint32 `json:"f1770"`
	F1771               uint32 `json:"f1771"`
	F1772               uint32 `json:"f1772"`
	F1773               uint32 `json:"f1773"`
	F1774               uint32 `json:"f1774"`
	F21                 uint32 `json:"f21"`
	F22                 uint32 `json:"f22"`
	F23                 uint32 `json:"f23"`
	F24                 uint32 `json:"f24"`
	F25                 uint32 `json:"f25"`
	F26                 uint32 `json:"f26"`
	F27                 uint32 `json:"f27"`
	F301                uint32 `json:"f301"`
	F302                uint32 `json:"f302"`
	F303                uint32 `json:"f303"`
	F401                uint32 `json:"f401"`
	F402                uint32 `json:"f402"`
	F403                uint32 `json:"f403"`
	F404                uint32 `json:"f404"`
	F405                uint32 `json:"f405"`
	F406                uint32 `json:"f406"`
	F407                uint32 `json:"f407"`
	F408                uint32 `json:"f408"`
	F409                uint32 `json:"f409"`
	F410                uint32 `json:"f410"`
	F411                uint32 `json:"f411"`
	F412                uint32 `json:"f412"`
	F413                uint32 `json:"f413"`
	F414                uint32 `json:"f414"`
	F415                uint32 `json:"f415"`
	F416                string `json:"f416"`
	F6                  uint32 `json:"f6"`
	F7                  uint32 `json:"f7"`
	F7a                 uint32 `json:"f7a"`
	F1001               uint32 `json:"f1001"`
	F1002               string `json:"f1002"`
	F1601               uint32 `json:"f1601"`
	F1602               uint32 `json:"f1602"`
	F1603               uint32 `json:"f1603"`
	F1604               uint32 `json:"f1604"`
	F1605               uint32 `json:"f1605"`
	F1606               uint32 `json:"f1606"`
	F1607               uint32 `json:"f1607"`
	F1608               uint32 `json:"f1608"`
	F1609               uint32 `json:"f1609"`
	F1610               uint32 `json:"f1610"`
	F1611               uint32 `json:"f1611"`
	F1612               uint32 `json:"f1612"`
	F1613               uint32 `json:"f1613"`
	F1614               string `json:"f1614"`
	NamaAtasan          string `json:"nama_atasan"`
	HpAtasan            string `json:"hp_atasan"`
	EmailAtasan         string `json:"email_atasan"`
	TinggalSelamaKuliah string `json:"tinggal_selama_kuliah"`
}

func UpdatePKTS(ctx *gin.Context, c pb.PKTSServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	nim := ctx.Param("nim")

	b := UpdatePKTSRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.UpdatePKTS(grpcCtx, &pb.PKTS{
		Nim:                 nim,
		F8:                  b.F8,
		F504:                b.F504,
		F502:                b.F502,
		F506:                b.F506,
		F505:                b.F505,
		F5A1:                b.F5a1,
		F5A2:                b.F5a2,
		F1101:               b.F1101,
		F1102:               b.F1102,
		F5B:                 b.F5b,
		F5C:                 b.F5c,
		F5D:                 b.F5d,
		F18A:                b.F18a,
		F18B:                b.F18b,
		F18C:                b.F18c,
		F18D:                b.F18d,
		F1201:               b.F1201,
		F1202:               b.F1202,
		F14:                 b.F14,
		F15:                 b.F15,
		F1761:               b.F1761,
		F1762:               b.F1762,
		F1763:               b.F1763,
		F1764:               b.F1764,
		F1765:               b.F1765,
		F1766:               b.F1766,
		F1767:               b.F1767,
		F1768:               b.F1768,
		F1769:               b.F1769,
		F1770:               b.F1770,
		F1771:               b.F1771,
		F1772:               b.F1772,
		F1773:               b.F1773,
		F1774:               b.F1774,
		F21:                 b.F21,
		F22:                 b.F22,
		F23:                 b.F23,
		F24:                 b.F24,
		F25:                 b.F25,
		F26:                 b.F26,
		F27:                 b.F27,
		F301:                b.F301,
		F302:                b.F302,
		F303:                b.F303,
		F401:                b.F401,
		F402:                b.F402,
		F403:                b.F403,
		F404:                b.F404,
		F405:                b.F405,
		F406:                b.F406,
		F407:                b.F407,
		F408:                b.F408,
		F409:                b.F409,
		F410:                b.F410,
		F411:                b.F411,
		F412:                b.F412,
		F413:                b.F413,
		F414:                b.F414,
		F415:                b.F415,
		F416:                b.F416,
		F6:                  b.F6,
		F7:                  b.F7,
		F7A:                 b.F7a,
		F1001:               b.F1001,
		F1002:               b.F1002,
		F1601:               b.F1601,
		F1602:               b.F1602,
		F1603:               b.F1603,
		F1604:               b.F1604,
		F1605:               b.F1605,
		F1606:               b.F1606,
		F1607:               b.F1607,
		F1608:               b.F1608,
		F1609:               b.F1609,
		F1610:               b.F1610,
		F1611:               b.F1611,
		F1612:               b.F1612,
		F1613:               b.F1613,
		F1614:               b.F1614,
		NamaAtasan:          b.NamaAtasan,
		HpAtasan:            b.HpAtasan,
		EmailAtasan:         b.EmailAtasan,
		TinggalSelamaKuliah: b.TinggalSelamaKuliah,
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
