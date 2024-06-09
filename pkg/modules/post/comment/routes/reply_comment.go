package routes

import (
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ReplyCommentRequestBody struct {
	CommentId uint64 `json:"comment_id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
}

func ReplyComment(ctx *gin.Context, c pb.CommentServiceClient) {
	var reqBody ReplyCommentRequestBody
	err := ctx.BindJSON(&reqBody)

	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to bind request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.ReplyComment(ctx, &pb.Comment{
		CommentId: reqBody.CommentId,
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
