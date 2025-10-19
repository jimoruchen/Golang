package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"fmt"
)

func oneToOne() {
	err := global.DB.Create(&models.UserModel{
		Name: "张三",
		Age:  18,
		UserDetailModel: &models.UserDetailModel{
			Email: "111@qq.com",
		},
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	global.Connect()
	//global.Migrate()
	oneToOne()
}
