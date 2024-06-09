package post

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/post/post/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *PostServiceClient {
	svc := &PostServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/posts")
	routes.GET("/", svc.GetAllPosts)
	routes.GET("/:id", svc.GetPostById)
	routes.DELETE("/:id", svc.DeletePost)

	return svc
}

func (svc *PostServiceClient) GetAllPosts(ctx *gin.Context) {
	routes.GetAllPosts(ctx, svc.Client)
}

func (svc *PostServiceClient) GetPostById(ctx *gin.Context) {
	routes.GetPostById(ctx, svc.Client)
}

func (svc *PostServiceClient) DeletePost(ctx *gin.Context) {
	routes.DeletePost(ctx, svc.Client)
}