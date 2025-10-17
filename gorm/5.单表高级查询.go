package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"fmt"

	"gorm.io/gorm"
)

func highQuery() {
	global.DB = global.DB.Debug()

	//批量插入
	//userList := []models.UserModel{
	//	{Name: "张三", Age: 20},
	//	{Name: "李四", Age: 20},
	//	{Name: "王五", Age: 18},
	//	{Name: "Alex", Age: 18},
	//	{Name: "Bob", Age: 21},
	//}
	//for _, user := range userList {
	//	err := global.DB.Create(&user).Error
	//	if err != nil {
	//		fmt.Printf("create user err:%v\n", err)
	//		continue
	//	}
	//	fmt.Printf("成功，user:%#v\n", user.Name)
	//}

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

	//分页
	//PaginateUsers(global.DB, 3, 2)

	//Scope
	//var users []models.UserModel
	//global.DB.Scopes(Age18).Find(&users)
	//fmt.Println(users)

	//var users []models.UserModel
	//global.DB.Scopes(NameIn("张三", "李四")).Find(&users)
	//fmt.Println(users)

	//原生SQL
	type User struct {
		Name string
		Age  int
	}
	var user []User
	global.DB.Raw("select name, age from user_models").Scan(&user)
	fmt.Println(user)
	global.DB.Exec("update user_models set age = 22 where id = 1")
}

func NameIn(NameList ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name in (?)", NameList)
	}
}

func Age18(tx *gorm.DB) *gorm.DB {
	return tx.Where("age > ?", 18)
}

func PaginateUsers(db *gorm.DB, pageNum, pageSize int) {
	var users []models.UserModel
	var count int64
	db.Model(&models.UserModel{}).Count(&count)
	offset := (pageNum - 1) * pageSize
	err := db.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("count: %d\n", count)
	fmt.Printf("%#v\n", users)
}

func main() {
	global.Connect()
	highQuery()
}
