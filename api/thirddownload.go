package api

import (
	"breath-server/model"
	"breath-server/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type MobanReq struct {
	Url    string `json:"url"`
	Serect string `json:"serect"`
}
type MubanDownloadInfo struct {
	NeedCoin  int    `json:"needCoin"`
	UserCoin  int    `json:"userCoin"`
	IsBuy     bool   `json:"isbuy"`
	MobanName string `json:"mobanname"`
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Data      string `json:"data"`
}

func DownloadMoban(c *gin.Context) {
	mobanReq := MobanReq{}
	c.BindJSON(&mobanReq)
	cardnumber := model.SelectBySerect(mobanReq.Serect)
	log.Printf("%v", cardnumber)
	if cardnumber.Id == 0 {
		c.JSON(901, gin.H{"msg": "该卡密不存在，请先购买卡密"})
		return
	}
	if cardnumber.Status == 1 {
		c.JSON(901, gin.H{"msg": "该卡密已经使用，不可重复使用"})
		return
	}
	//模版之家id解析
	end := strings.Split(mobanReq.Url, "/")
	id := strings.Split(end[len(end)-1], ".")[0]
	payload := "userid=596116&mobanid=" + id + "&screkey=b6d617e284370fa51d2f2e64cd773a6a"
	mubanDownloadInfo := MubanDownloadInfo{}
	if err := json.Unmarshal(utils.DoPostFrom("http://vip.cssmoban.com/ApiShenJi/down", payload), &mubanDownloadInfo); err != nil {
		c.JSON(901, gin.H{"msg": "下载出错了"})
		return

	}
	c.JSON(200, gin.H{"msg": "该卡密已经使用，不可重复使用", "data": mubanDownloadInfo})
	return

}
