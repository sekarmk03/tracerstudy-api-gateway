package main

import (
	"log"
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/auth/auth"
	"tracerstudy-api-gateway/pkg/modules/auth/user"
	"tracerstudy-api-gateway/pkg/modules/tracer/kabkota"
	"tracerstudy-api-gateway/pkg/modules/tracer/provinsi"
	"tracerstudy-api-gateway/pkg/modules/tracer/pkts"
	"tracerstudy-api-gateway/pkg/modules/tracer/prodi"
	"tracerstudy-api-gateway/pkg/modules/tracer/responden"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Could not load config:", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)
	user.RegisterRoutes(r, &c)
	kabkota.RegisterRoutes(r, &c)
	provinsi.RegisterRoutes(r, &c)
	pkts.RegisterRoutes(r, &c)
	prodi.RegisterRoutes(r, &c)
	responden.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
