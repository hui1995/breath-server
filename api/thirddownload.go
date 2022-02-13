package api

import (
	"breath-server/model"
	"github.com/gin-gonic/gin"
	"log"
)

type MobanReq struct {
	Url    string `json:"url"`
	Serect string `json:"serect"`
}
type FininshReq struct {
	Id string `json:"id"`
}

func DownloadMoban(c *gin.Context) {
	mobanReq := MobanReq{}
	c.BindJSON(&mobanReq)
	cardnumber := model.SelectBySerect(mobanReq.Serect)
	log.Printf("%v", cardnumber)
	if cardnumber.Id == 0 {
		c.JSON(200, gin.H{"code": 50, "msg": "该卡密不存在，请先购买卡密"})
		return
	}
	if cardnumber.Status == 1 {
		c.JSON(200, gin.H{"code": 50, "msg": "该卡密已经使用，不可重复使用"})
		return
	}
	//serect,访问携带数据:
	serect := cardnumber.Id
	c.JSON(200, gin.H{"code": 200, "msg": "ok", "data": serect})
	return

}
func DownloadMobanFinsh(c *gin.Context) {
	finshReq := FininshReq{}
	c.BindJSON(&finshReq)
	model.ChangeStatus(1)
	c.JSON(200, gin.H{"code": 200, "msg": "ok"})
	return

}

//docker commit -a "bethmeta.com" -m "fronetnd" 440e96cb847f  breath-fronetnd
