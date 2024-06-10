package userstudy

import (
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/tracer/userstudy/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *UserStudyServiceClient {
	svc := &UserStudyServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/userstudy")
	routes.GET("/", svc.GetAllUserStudy)
	routes.POST("/:nim", svc.GetUserStudyByNim)
	routes.POST("/", svc.CreateUserStudy)
	routes.POST("/alumni", svc.GetAlumniListByAtasan)
	routes.PUT("/:nim", svc.UpdateUserStudy)
	routes.GET("/export", svc.ExportUSReport)
	routes.GET("/rekap", svc.GetUserStudyRekap)
	routes.GET("/rekap/:kode", svc.GetUserStudyRekapByProdi)

	return svc
}

func (svc *UserStudyServiceClient) GetAllUserStudy(ctx *gin.Context) {
	routes.GetAllUserStudy(ctx, svc.Client)
}

func (svc *UserStudyServiceClient) GetUserStudyByNim(ctx *gin.Context) {
	routes.GetUserStudyByNim(ctx, svc.Client)
}

func (svc *UserStudyServiceClient) CreateUserStudy(ctx *gin.Context) {
	routes.CreateUserStudy(ctx, svc.Client)
}

func (svc *UserStudyServiceClient) UpdateUserStudy(ctx *gin.Context) {
	routes.UpdateUserStudy(ctx, svc.Client)
}

func (svc *UserStudyServiceClient) ExportUSReport(ctx *gin.Context) {
	routes.ExportUSReport(ctx, svc.Client)
}

func (svc *UserStudyServiceClient) GetUserStudyRekap(ctx *gin.Context) {
	routes.GetUserStudyRekap(ctx, svc.Client)
}

func (svc *UserStudyServiceClient) GetUserStudyRekapByProdi(ctx *gin.Context) {
	routes.GetUserStudyRekapByProdi(ctx, svc.Client)
}

func (svc *UserStudyServiceClient) GetAlumniListByAtasan(ctx *gin.Context) {
	routes.GetAlumniListByAtasan(ctx, svc.Client)
}
