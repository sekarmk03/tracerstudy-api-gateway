package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type CreateUserRequestBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   uint32 `json:"role_id"`
}

func CreateUser(ctx *gin.Context, c pb.UserServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))
	
	b := CreateUserRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.CreateUser(grpcCtx, &pb.User{
		Name:     b.Name,
		Username: b.Username,
		Email:    b.Email,
		Password: b.Password,
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

	ctx.JSON(http.StatusCreated, &res)
}
