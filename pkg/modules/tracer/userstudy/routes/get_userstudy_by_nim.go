package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type GetUserStudyByNimRequest struct {
	Nim            string `json:"nim"`
	EmailResponden string `json:"email_responden"`
	HpResponden    string `json:"hp_responden"`
}

func GetUserStudyByNim(ctx *gin.Context, c pb.UserStudyServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	b := GetUserStudyByNimRequest{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.GetUserStudyByNim(grpcCtx, &pb.GetUserStudyByNimRequest{
		Nim:            b.Nim,
		EmailResponden: b.EmailResponden,
		HpResponden:    b.HpResponden,
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
