package router

import (
	"5/exam/nacos/handler/service"
	"5/exam/nacos/util"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)
	api := r.Use(util.Cors())
	{
		api.GET("/BannerData", service.BannerData)
	}
	return r
}
