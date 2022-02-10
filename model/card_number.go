package model

type CardNumber struct {
	Id          int    `gorm:"column:id"`
	Serect      string `gorm:"column:serect"`
	Status      int    `gorm:"column:status"`
	Legal       int    `gorm:"column:legal"`
	Type        int    `gorm:"column:type"`
	CardPowerId int    `gorm:"column:power_card_id"`
	CreateDate  string `gorm:"column:create_date"`
}

// 设置表名
func (CardNumber) TableName() string {
	return "card_number"
}

// 添加一行数据
//func AddMember() {
//	member := CardNumber{Name: "zhangsan", Age: 18, Sex: 1}
//
//	result := Db.Create(&member)
//	fmt.Println(member.Id)
//	fmt.Println(result.RowsAffected)
//	fmt.Println(result.Error)
//
//}

func SelectBySerect(serect string) CardNumber {
	cardNumber := new(CardNumber)
	Db.Where(&CardNumber{Serect: serect}).First(&cardNumber)
	//defer Db.Close()
	return *cardNumber
}
