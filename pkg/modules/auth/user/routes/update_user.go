package routes

import (
	"context"
	"net/http"
	"strconv"
	"tracerstudy-api-gateway/pkg/pb"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   uint32 `json:"role_id"`
}

func UpdateUser(ctx *gin.Context, c pb.UserServiceClient) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	b := UpdateUserRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateUser(context.Background(), &pb.User{
		Id:       id,
		Username: b.Username,
		Email:    b.Email,
		RoleId:   b.RoleId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
