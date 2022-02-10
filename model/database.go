package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//https://learnku.com/docs/gorm/v1/query/3786#d765a4
// 公共Db对象
var Db *gorm.DB"db.port")
	password := viper.GetString("db.password")
	charset := viper.GetString("db.charset")
	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	Db = db
}
