package auth

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/auth/auth/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *AuthServiceClient {
	svc := &AuthServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.RegisterUser)
	routes.POST("/login", svc.LoginUser)
	routes.POST("/login-alumni", svc.LoginAlumni)
	routes.POST("/login-userstudy", svc.LoginUserStudy)
	routes.GET("/whoami", svc.GetCurrentUser)

	return svc
}

func (svc *AuthServiceClient) RegisterUser(ctx *gin.Context) {
	routes.RegisterUser(ctx, svc.Client)
}

func (svc *AuthServiceClient) LoginUser(ctx *gin.Context) {
	routes.LoginUser(ctx, svc.Client)
}

func (svc *AuthServiceClient) LoginAlumni(ctx *gin.Context) {
	routes.LoginAlumni(ctx, svc.Client)
}

func (svc *AuthServiceClient) LoginUserStudy(ctx *gin.Context) {
	routes.LoginUserStudy(ctx, svc.Client)
}

func (svc *AuthServiceClient) GetCurrentUser(ctx *gin.Context) {
	routes.GetCurrentUser(ctx, svc.Client)
}
