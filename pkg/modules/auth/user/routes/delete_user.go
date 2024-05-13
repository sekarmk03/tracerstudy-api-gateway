package routes

import (
	"context"
	"net/http"
	"strconv"
	"tracerstudy-api-gateway/pkg/pb"

	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context, c pb.UserServiceClient) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.DeleteUser(context.Background(), &pb.GetUserByIdRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
