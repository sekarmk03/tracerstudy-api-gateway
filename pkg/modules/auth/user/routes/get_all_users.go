package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetAllUsers(ctx *gin.Context, c pb.UserServiceClient) {
	res, err := c.GetAllUsers(context.Background(), &emptypb.Empty{})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}