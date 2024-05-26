package kabkota

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/tracer/kabkota/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *KabKotaServiceClient {
	svc := &KabKotaServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/kabkota")
	routes.GET("/", svc.GetAllKabKota)
	routes.GET("/:id", svc.GetKabKotaByIdWil)
	routes.POST("/", svc.CreateKabKota)
	routes.PUT("/:id", svc.UpdateKabKota)
	routes.DELETE("/:id", svc.DeleteKabKota)
	routes.GET("/provinsi/:provinsiId", svc.GetKabKotaByProvinsi)

	return svc
}

func (svc *KabKotaServiceClient) GetAllKabKota(ctx *gin.Context) {
	routes.GetAllKabKota(ctx, svc.Client)
}

func (svc *KabKotaServiceClient) GetKabKotaByProvinsi(ctx *gin.Context) {
	routes.GetKabKotaByProvinsi(ctx, svc.Client)
}

func (svc *KabKotaServiceClient) GetKabKotaByIdWil(ctx *gin.Context) {
	routes.GetKabKotaByIdWil(ctx, svc.Client)
}

func (svc *KabKotaServiceClient) CreateKabKota(ctx *gin.Context) {
	routes.CreateKabKota(ctx, svc.Client)
}

func (svc *KabKotaServiceClient) UpdateKabKota(ctx *gin.Context) {
	routes.UpdateKabKota(ctx, svc.Client)
}

func (svc *KabKotaServiceClient) DeleteKabKota(ctx *gin.Context) {
	routes.DeleteKabKota(ctx, svc.Client)
}
