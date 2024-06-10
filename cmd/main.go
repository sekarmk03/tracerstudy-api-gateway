package main

import (
	"log"
	"time"
	"tracerstudy-api-gateway/pkg/config"
	"tracerstudy-api-gateway/pkg/modules/auth/auth"
	"tracerstudy-api-gateway/pkg/modules/auth/user"
	"tracerstudy-api-gateway/pkg/modules/post/comment"
	"tracerstudy-api-gateway/pkg/modules/post/post"
	"tracerstudy-api-gateway/pkg/modules/tracer/kabkota"
	"tracerstudy-api-gateway/pkg/modules/tracer/pkts"
	"tracerstudy-api-gateway/pkg/modules/tracer/prodi"
	"tracerstudy-api-gateway/pkg/modules/tracer/provinsi"
	"tracerstudy-api-gateway/pkg/modules/tracer/responden"
	"tracerstudy-api-gateway/pkg/modules/tracer/userstudy"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Could not load config:", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth.RegisterRoutes(r, &c)
	user.RegisterRoutes(r, &c)
	kabkota.RegisterRoutes(r, &c)
	provinsi.RegisterRoutes(r, &c)
	pkts.RegisterRoutes(r, &c)
	prodi.RegisterRoutes(r, &c)
	responden.RegisterRoutes(r, &c)
	userstudy.RegisterRoutes(r, &c)
	post.RegisterRoutes(r, &c)
	comment.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
