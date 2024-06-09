package routes

import (
	"net/http"
	"strconv"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ReplyCommentRequestBody struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func ReplyComment(ctx *gin.Context, c pb.CommentServiceClient) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to convert id to int")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	var reqBody ReplyCommentRequestBody
	err = ctx.BindJSON(&reqBody)

	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to bind request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.ReplyComment(ctx, &pb.Comment{
		CommentId: id,
		Name:      reqBody.Name,
		Content:   reqBody.Content,
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
