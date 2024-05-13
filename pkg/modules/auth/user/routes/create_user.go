package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"

	"github.com/gin-gonic/gin"
)

type CreateUserRequestBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   uint32 `json:"role_id"`
}

func CreateUser(ctx *gin.Context, c pb.UserServiceClient) {
	b := CreateUserRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateUser(context.Background(), &pb.User{
		Name:     b.Name,
		Username: b.Username,
		Email:    b.Email,
		Password: b.Password,
		RoleId:   b.RoleId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
