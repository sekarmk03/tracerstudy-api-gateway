package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LoginAlumniRequestBody struct {
	Nim           string `json:"nim"`
	TanggalSidang string `json:"tanggal_sidang"`
}

func LoginAlumni(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginAlumniRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.LoginAlumni(context.Background(), &pb.LoginAlumniRequest{
		Nim:           b.Nim,
		TanggalSidang: b.TanggalSidang,
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
