package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"fmt"
)

func UserZdy() {
	//创建
	//err := global.DB.Create(&models.UserZdy{
	//	Name: "张三",
	//	Info: models.UserInfo{
	//		Likes: []string{"唱", "跳", "rap"},
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//查询
	//var userZdy models.UserZdy
	//global.DB.Take(&userZdy)
	//fmt.Println(userZdy, userZdy.Info.Likes)

	//创建
	//err := global.DB.Create(&models.UserZdy{
	//	Name: "张三",
	//	Info: models.UserInfo{
	//		Likes: []string{"唱", "跳", "rap"},
	//	},
	//	CardInfo: models.Card{
	//		Card: []string{"法拉利", "兰博基尼"},
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//查询
	var userZdy models.UserZdy
	global.DB.Last(&userZdy)
	fmt.Println(userZdy, userZdy.Info.Likes, userZdy.CardInfo.Card)
}

func main() {
	global.Connect()
	//global.Migrate()
	UserZdy()
}
