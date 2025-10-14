package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"fmt"
)

func highQuery() {
	global.DB = global.DB.Debug()

	//var user models.UserModel
	//global.DB.Where("age > ?", 18).Take(&user)
	//fmt.Println(user)

	//结构体
	//var user models.UserModel
	//global.DB.Where(models.UserModel{
	//	Name: "张三",
	//}).Take(&user)
	//fmt.Println(user)

	//map
	//var user models.UserModel
	//global.DB.Where(map[string]any{
	//	"name": "张三",
	//}).Take(&user)
	//fmt.Println(user)

	//嵌套
	//var user models.UserModel
	//query := global.DB.Where("age > 18")
	//global.DB.Where(query).Take(&user)
	//fmt.Println(user)

	//Or
	//var user []models.UserModel
	//global.DB.Or("name = ?", "张三").Or("age = ?", 18).Find(&user)
	//fmt.Println(user)

	//Not
	//var user []models.UserModel
	//global.DB.Not("age = ?", 18).Find(&user)
	//fmt.Println(user)

	//Order
	var user []models.UserModel
	global.DB.Order("age desc").Order("id asc").Find(&user)
	fmt.Println(user)

}

func main() {
	global.Connect()
	highQuery()
}
