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

func DownloadMoban(c *gin.Context) {
	mobanReq := MobanReq{}
	c.BindJSON(&mobanReq)
	cardnumber := model.SelectBySerect(mobanReq.Serect)
	log.Printf("%v", cardnumber)
	if cardnumber.Id == 0 {
		c.JSON(201, gin.H{"msg": "该卡密不存在，请先购买卡密"})
		return
	}
	if cardnumber.Status == 1 {
		c.JSON(201, gin.H{"msg": "该卡密已经使用，不可重复使用"})
		return
	}
	//sourceInfo:=model.SelectBySourceUrl(mobanReq.Url)

	log.Printf("%v", cardnumber.Id)

	c.JSON(200, gin.H{"msg": "Hello Go"})
}
