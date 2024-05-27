package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type UpdateRespondenRequestBody struct {
	Hp            string `json:"hp"`
	Nik           string `json:"nik"`
	Npwp          string `json:"npwp"`
	Email         string `json:"email"`
	TanggalWisuda string `json:"tanggal_wisuda"`
}

func UpdateResponden(ctx *gin.Context, c pb.RespondenServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	nim := ctx.Param("nim")

	b := UpdateRespondenRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.UpdateResponden(grpcCtx, &pb.Responden{
		Nim:           nim,
		Hp:            b.Hp,
		Nik:           b.Nik,
		Npwp:          b.Npwp,
		Email:         b.Email,
		TanggalWisuda: b.TanggalWisuda,
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
