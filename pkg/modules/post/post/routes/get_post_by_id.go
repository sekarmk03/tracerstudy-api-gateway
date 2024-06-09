package routes

import (
	"net/http"
	"strconv"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

func GetPostById(ctx *gin.Context, c pb.PostServiceClient) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to convert id to int")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.GetPostById(ctx, &pb.GetPostByIdRequest{
		Id: id,
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
