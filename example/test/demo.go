package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/v2"
	"github.com/xgpc/dsg/v2/models/cond"
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

	var list []User
	err = dsg.DB().Scopes(cond.Page(1, 10)).Find(&list).Error
	if err != nil {
		panic(err)
	}

	var ctx iris.Context
	dsg.DB().Scopes(cond.PageByQuery(ctx)).Find(&list)

	// cond.Eq id = 1
	dsg.DB().Scopes(cond.Eq("id", 1)).Find(&list)
	// cond.NotEq id != 1
	dsg.DB().Scopes(cond.NotEq("id", 1)).Find(&list)
	// cond.Gt id > 1
	dsg.DB().Scopes(cond.Gt("id", 1)).Find(&list)
	// cond.Gte id >= 1
	dsg.DB().Scopes(cond.Gte("id", 1)).Find(&list)
	// cond.Lt id < 1
	dsg.DB().Scopes(cond.Lt("id", 1)).Find(&list)
	// cond.Lte id <= 1
	dsg.DB().Scopes(cond.Lte("id", 1)).Find(&list)
	// cond.Like name like '%test%'
	dsg.DB().Scopes(cond.Like("name", "test")).Find(&list)
	// cond.Starting name like 'test%'
	dsg.DB().Scopes(cond.Starting("name", "test")).Find(&list)
	// cond.Ending name like '%test'
	dsg.DB().Scopes(cond.Ending("name", "test")).Find(&list)
	// cond.In id in (1,2,3)
	dsg.DB().Scopes(cond.In("id", []int{1, 2, 3})).Find(&list)
	// cond.NotIn  id not in (1,2,3)
	dsg.DB().Scopes(cond.NotIn("id", []int{1, 2, 3})).Find(&list)
	// cond.Between id between 1 and 10
	dsg.DB().Scopes(cond.Between("id", 1, 10)).Find(&list)
}

type User struct {
	ID   int64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

func (User) TableName() string {
	return "user"
}
