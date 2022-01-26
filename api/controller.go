package api

import "github.com/gin-gonic/gin"

// 首页控制器
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Hello Go"})
}

//// 测试控制器
//func PostTest(c *gin.Context) {
//
//	//data := make(map[string]interface{})
//	//
//	//data["foo"] = "bar"
//	//fmt.Println(data)
//	//c.JSON(200, data)
//	//common.CommonFail(c, rest, "失败了")
//}
//
//// 测试数据库操作
//func GetTestSql(c *gin.Context) {
//
//	//model.AddMember()
//	//fmt.Println("执行完成")
//
//	fmt.Println(viper.GetString("wx.appid"))
//
//}
