package route

import (
	"breath-server/api"
	"breath-server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Cors())
	download := router.Group("/download")
	{
		download.POST("/moban", api.DownloadMoban)
		download.POST("/mobanfininsh", api.DownloadMobanFinsh)

	}

	// 首页路由
	router.GET("/index", api.GetIndex)

	return router
}
