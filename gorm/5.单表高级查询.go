package main

import (
	"Golang/gorm/global"
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
	//var user []models.UserModel
	//global.DB.Order("age desc").Order("id asc").Find(&user)
	//fmt.Println(user)

	//Scan 只需要特定字段
	//var nameList []string
	//var user models.UserModel
	//global.DB.Model(user).Select("name").Scan(&nameList)
	//global.DB.Model(models.UserModel{}).Select("name").Scan(&nameList)
	//global.DB.Model(models.UserModel{}).Pluck("name", &nameList)
	//fmt.Println(nameList)

	//只需要几个特定的字段
	//type User struct {
	//	Name string
	//	Age  int
	//}
	//type User struct {
	//	Label string `gorm:"column:name"` //首字母必须大写
	//	Value int    `gorm:"column:age"`
	//}
	//
	//var userList []User
	//global.DB.Model(models.UserModel{}).Scan(&userList)
	//fmt.Println(userList)

	//分组
	//type User struct {
	//	Age   int
	//	Count int
	//}
	//var user []User
	//global.DB.Model(models.UserModel{}).
	//	Group("age").
	//	Select("age, count(*) as count").
	//	Scan(&user)
	//fmt.Println(user)		//[{22 1} {23 1} {18 2} {17 1}]

	//去重
	//var ageList []int
	//global.DB.Model(models.UserModel{}).Distinct("age").Pluck("age", &ageList)
	//fmt.Println(ageList)

}

func main() {
	global.Connect()
	highQuery()
}
