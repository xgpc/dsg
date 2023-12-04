package main

import (
	"fmt"
	"github.com/xgpc/dsg/v2"
	"github.com/xgpc/dsg/v2/pkg/util"
)

func main() {
	dsg.Load("config.yml")
	fmt.Println(dsg.Conf)

	userID := 1

	// 查
	var user User
	err := dsg.DB().Model(user).First(&user, userID).Error
	if err != nil {
		panic(err)
	}

	// 增
	user = User{
		Name: "test",
	}
	err = dsg.DB().Model(user).Create(&user).Error
	if err != nil {
		panic(err)
	}

	// 删
	err = dsg.DB().Model(user).Delete(&user, userID).Error
	if err != nil {
		panic(err)
	}

	body := util.StructToMap(user)
	// 改
	err = dsg.DB().Model(user).Where("id", userID).Updates(&body).Error
	if err != nil {
		panic(err)
	}
}

type User struct {
	ID   int64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

func (User) TableName() string {
	return "user"
}
