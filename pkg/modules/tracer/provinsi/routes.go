package provinsi

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/tracer/provinsi/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ProvinsiServiceClient {
	svc := &ProvinsiServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/provinsi")
	routes.GET("/", svc.GetAllProvinsi)
	routes.GET("/:id", svc.GetProvinsiByIdWil)
	routes.POST("/", svc.CreateProvinsi)
	routes.PUT("/:id", svc.UpdateProvinsi)
	routes.DELETE("/:id", svc.DeleteProvinsi)

	return svc
}

func (svc *ProvinsiServiceClient) GetAllProvinsi(ctx *gin.Context) {
	routes.GetAllProvinsi(ctx, svc.Client)
}

func (svc *ProvinsiServiceClient) GetProvinsiByIdWil(ctx *gin.Context) {
	routes.GetProvinsiByIdWil(ctx, svc.Client)
}

func (svc *ProvinsiServiceClient) CreateProvinsi(ctx *gin.Context) {
	routes.CreateProvinsi(ctx, svc.Client)
}

func (svc *ProvinsiServiceClient) UpdateProvinsi(ctx *gin.Context) {
	routes.UpdateProvinsi(ctx, svc.Client)
}

func (svc *ProvinsiServiceClient) DeleteProvinsi(ctx *gin.Context) {
	routes.DeleteProvinsi(ctx, svc.Client)
}
