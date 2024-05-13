package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LoginUserStudyRequestBody struct {
	EmailAtasan string `json:"email_atasan"`
	HpAtasan    string `json:"hp_atasan"`
	NamaAtasan  string `json:"nama_atasan"`
}

func LoginUserStudy(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginUserStudyRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.LoginUserStudy(context.Background(), &pb.LoginUserStudyRequest{
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
