package routes

import (
	"context"
	"net/http"
	"strconv"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type UpdateUserRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   uint32 `json:"role_id"`
}

func UpdateUser(ctx *gin.Context, c pb.UserServiceClient) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to convert id to int")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	b := UpdateUserRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	res, err := c.UpdateUser(grpcCtx, &pb.User{
		Id:       id,
		Username: b.Username,
		Email:    b.Email,
		RoleId:   b.RoleId,
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
