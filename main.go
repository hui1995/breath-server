package main

import (
	"breath-server/config"
	"breath-server/model"
	"breath-server/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := route.InitRouter()

	//加载配置文件
	config.InitConfig()

	// 创建数据库连接
	model.InitDb()
	// defer关键字，当API访问结束时，关闭数据库连接
	defer model.Db.Close()

	//监听8000端口
	gin.SetMode(viper.GetString("server.run_mode"))
	r.Run(viper.GetString("server.addr"))
}
