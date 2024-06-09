package comment

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/post/comment/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *CommentServiceClient {
	svc := &CommentServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/comments")
	routes.GET("/", svc.GetAllComments)
	routes.GET("/:id", svc.GetCommentById)
	routes.GET("/post/:id", svc.GetCommentsByPost)
	routes.POST("/", svc.CreateComment)
	routes.POST("/:id/reply", svc.ReplyComment)
	routes.DELETE("/:id", svc.DeleteComment)

	return svc
}

func (svc *CommentServiceClient) GetAllComments(ctx *gin.Context) {
	routes.GetAllComments(ctx, svc.Client)
}

func (svc *CommentServiceClient) GetCommentById(ctx *gin.Context) {
	routes.GetCommentById(ctx, svc.Client)
}

func (svc *CommentServiceClient) CreateComment(ctx *gin.Context) {
	routes.CreateComment(ctx, svc.Client)
}

func (svc *CommentServiceClient) ReplyComment(ctx *gin.Context) {
	routes.ReplyComment(ctx, svc.Client)
}

func (svc *CommentServiceClient) GetCommentsByPost(ctx *gin.Context) {
	routes.GetCommentsByPost(ctx, svc.Client)
}

func (svc *CommentServiceClient) DeleteComment(ctx *gin.Context) {
	routes.DeleteComment(ctx, svc.Client)
}