package responden

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/tracer/responden/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *RespondenServiceClient {
	svc := &RespondenServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/responden")
	routes.GET("/", svc.GetAllResponden)
	routes.GET("/:nim", svc.GetRespondenByNim)
	routes.GET("/siak/:nim", svc.UpdateRespondenFromSiak)
	routes.POST("/", svc.CreateResponden)
	routes.PUT("/:nim", svc.UpdateResponden)
	routes.POST("/list", svc.GetRespondenByNimList)

	return svc
}

func (svc *RespondenServiceClient) GetAllResponden(ctx *gin.Context) {
	routes.GetAllResponden(ctx, svc.Client)
}

func (svc *RespondenServiceClient) GetRespondenByNim(ctx *gin.Context) {
	routes.GetRespondenByNim(ctx, svc.Client)
}

func (svc *RespondenServiceClient) UpdateRespondenFromSiak(ctx *gin.Context) {
	routes.UpdateRespondenFromSiak(ctx, svc.Client)
}

func (svc *RespondenServiceClient) CreateResponden(ctx *gin.Context) {
	routes.CreateResponden(ctx, svc.Client)
}

func (svc *RespondenServiceClient) UpdateResponden(ctx *gin.Context) {
	routes.UpdateResponden(ctx, svc.Client)
}

func (svc *RespondenServiceClient) GetRespondenByNimList(ctx *gin.Context) {
	routes.GetRespondenByNimList(ctx, svc.Client)
}

func (svc *RespondenServiceClient) GetOrCreateResponden(ctx *gin.Context) {
	routes.GetOrCreateResponden(ctx, svc.Client)
}
