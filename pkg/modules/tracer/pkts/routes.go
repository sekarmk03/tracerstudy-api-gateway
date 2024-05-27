package pkts

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/tracer/pkts/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *PKTSServiceClient {
	svc := &PKTSServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/pkts")
	routes.GET("/", svc.GetAllPKTS)
	routes.GET("/:nim", svc.GetPKTSByNim)
	routes.GET("/rekap/:tahunSidang", svc.GetPKTSRekapByYear)
	routes.GET("/rekap/:tahunSidang/:kode", svc.GetPKTSRekapByProdi)
	routes.POST("/", svc.CreatePKTS)
	routes.PUT("/:nim", svc.UpdatePKTS)
	routes.GET("/export/:tahun", svc.ExportPKTSReport)

	return svc
}

func (svc *PKTSServiceClient) GetAllPKTS(ctx *gin.Context) {
	routes.GetAllPKTS(ctx, svc.Client)
}

func (svc *PKTSServiceClient) GetPKTSByNim(ctx *gin.Context) {
	routes.GetPKTSByNim(ctx, svc.Client)
}

func (svc *PKTSServiceClient) CreatePKTS(ctx *gin.Context) {
	routes.CreatePKTS(ctx, svc.Client)
}

func (svc *PKTSServiceClient) UpdatePKTS(ctx *gin.Context) {
	routes.UpdatePKTS(ctx, svc.Client)
}

func (svc *PKTSServiceClient) ExportPKTSReport(ctx *gin.Context) {
	routes.ExportPKTSReport(ctx, svc.Client)
}

func (svc *PKTSServiceClient) GetPKTSRekapByYear(ctx *gin.Context) {
	routes.GetPKTSRekapByYear(ctx, svc.Client)
}

func (svc *PKTSServiceClient) GetPKTSRekapByProdi(ctx *gin.Context) {
	routes.GetPKTSRekapByProdi(ctx, svc.Client)
}
