package prodi

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/tracer/prodi/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ProdiServiceClient {
	svc := &ProdiServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/prodi")
	routes.GET("/", svc.GetAllProdi)
	routes.GET("/fakultas", svc.GetAllFakultas)
	routes.GET("/:kode", svc.GetProdiByKode)
	routes.POST("/", svc.CreateProdi)
	routes.PUT("/:kode", svc.UpdateProdi)
	routes.DELETE("/:kode", svc.DeleteProdi)

	return svc
}

func (svc *ProdiServiceClient) GetAllProdi(ctx *gin.Context) {
	routes.GetAllProdi(ctx, svc.Client)
}

func (svc *ProdiServiceClient) GetAllFakultas(ctx *gin.Context) {
	routes.GetAllFakultas(ctx, svc.Client)
}

func (svc *ProdiServiceClient) GetProdiByKode(ctx *gin.Context) {
	routes.GetProdiByKode(ctx, svc.Client)
}

func (svc *ProdiServiceClient) CreateProdi(ctx *gin.Context) {
	routes.CreateProdi(ctx, svc.Client)
}

func (svc *ProdiServiceClient) UpdateProdi(ctx *gin.Context) {
	routes.UpdateProdi(ctx, svc.Client)
}

func (svc *ProdiServiceClient) DeleteProdi(ctx *gin.Context) {
	routes.DeleteProdi(ctx, svc.Client)
}
