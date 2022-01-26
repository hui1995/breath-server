package route

import (
	"breath-server/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 首页路由
	router.GET("/index", api.GetIndex)

	return router
}
