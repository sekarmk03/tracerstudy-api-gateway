package user

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/auth/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *UserServiceClient {
	svc := &UserServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/users")
	routes.GET("/", svc.GetAllUsers)
	routes.GET("/:id", svc.GetUserById)
	routes.POST("/", svc.CreateUser)
	routes.PUT("/:id", svc.UpdateUser)
	routes.DELETE("/:id", svc.DeleteUser)

	return svc
}

func (svc *UserServiceClient) GetAllUsers(ctx *gin.Context) {
	routes.GetAllUsers(ctx, svc.Client)
}

func (svc *UserServiceClient) GetUserById(ctx *gin.Context) {
	routes.GetUserById(ctx, svc.Client)
}

func (svc *UserServiceClient) CreateUser(ctx *gin.Context) {
	routes.CreateUser(ctx, svc.Client)
}

func (svc *UserServiceClient) UpdateUser(ctx *gin.Context) {
	routes.UpdateUser(ctx, svc.Client)
}

func (svc *UserServiceClient) DeleteUser(ctx *gin.Context) {
	routes.DeleteUser(ctx, svc.Client)
}
