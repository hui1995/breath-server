package model

import "fmt"

type Pw_member struct {
	Id   int
	Name string
	Age  int
	Sex  int
}

// 设置表名
func (Pw_member) TableName() string {
	return "pw_member"
}

// 添加一行数据
func AddMember() {
	member := Pw_member{Name: "zhangsan", Age: 18, Sex: 1}

	result := Db.Create(&member)
	fmt.Println(member.Id)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

}
