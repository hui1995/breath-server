package model

import "fmt"

type SourceInfo struct {
	Id         int    `gorm:"column:id"`
	Title      string `gorm:"column:title"`
	CateId     int    `gorm:"column:cate_id"`
	Images     string `gorm:"column:image"`
	Tags       string `gorm:"column:tags"`
	ZipUrl     string `gorm:"column:zip_url"`
	Content    string `gorm:"column:content"`
	IsDelete   int    `gorm:"column:is_delete"`
	Like       int    `gorm:"column:like"`
	Unlike     int    `gorm:"column:unlike"`
	Download   int    `gorm:"column:download"`
	View       int    `gorm:"column:view"`
	SourceId   string `gorm:"column:source_id"`
	SourceUrl  string `gorm:"column:source_url"`
	SourceWeb  string `gorm:"column:source_web"`
	sourceType string `gorm:"column:source_type"`
	CreateDate string `gorm:"column:create_date"`
}

// 设置表名
func (SourceInfo) TableName() string {
	return "source_info"
}

// 添加一行数据
func Add(sourceInfo SourceInfo) {
	//
	result := Db.Create(&sourceInfo)
	fmt.Println(result.RowsAffected)
	//	fmt.Println(result.Error)
	//
}

//拆根据url查询是否存在数据
func SelectBySourceUrl(href string) SourceInfo {
	sourceInfo := new(SourceInfo)
	Db.Where(&SourceInfo{SourceUrl: href}).First(&sourceInfo)
	//defer Db.Close()
	return *sourceInfo
}
func updateViewAndSource(zip_url string, id int) {
	var sourceInfo SourceInfo
	Db.First(&sourceInfo, id).Updates(SourceInfo{View: sourceInfo.View + 1, ZipUrl: zip_url})

}
