package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func insert() {
	//err := global.DB.Create(&models.UserModel{
	//	Name: "王五",
	//	Age:  18,
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//回填式创建
	//user := models.UserModel{
	//	Name: "李四",
	//	Age:  18,
	//}
	//err := global.DB.Create(&user).Error
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(user.Name, user.ID, user.Age, user.CreatedAt)

	//批量插入
	//userList := []models.UserModel{
	//	{Name: "Alex", Age: 18},
	//	{Name: "Tom", Age: 20},
	//	{Name: "Invalid", Age: -5}, // 假设这个会触发数据库约束失败
	//}
	//for _, user := range userList {
	//	err := global.DB.Create(&user).Error
	//	if err != nil {
	//		fmt.Printf("创建用户 %s 失败：%v", user.Name, err)
	//		continue
	//	}
	//	fmt.Printf(" 创建用户 %s 成功", user.Name)
	//}

	//钩子函数
	err := global.DB.Create(&models.UserModel{
		Name: "王五1",
		Age:  18,
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}

func query() {
	var userList []models.UserModel

	//全部查询
	//err := global.DB.Find(&userList).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(userList)

	//特定查询
	global.DB.Where("age >= ?", 20).Find(&userList)
	//global.DB.Find(&userList, "age >= ?", 20)
	fmt.Println(userList)

	var user models.UserModel
	//查一个
	//global.DB.Take(&user)

	//根据主键去查
	//global.DB.Take(&user, 2)

	//根据主键排序查第一个
	//global.DB.First(&user)

	//根据主键排序查最后一个，.Debug()会打印实际的SQL
	//global.DB.Debug().Last(&user)

	//user.ID = 3 //查询会携带主键
	//global.DB.Take(&user)

	//查一条没查到会抛出错误没查到
	//err := global.DB.Take(&user, 111).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	fmt.Println("记录不存在")
	//	return
	//}

	//使用Limit(1).Find()，不会抛出错误没查到
	err := global.DB.Limit(1).Find(&user, 111).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("记录不存在")
		return
	}
	fmt.Println(user)
}

func save() {
	global.DB = global.DB.Debug()
	var user = models.UserModel{
		ID:        2,
		Name:      "", //可以更新零值
		Age:       20,
		CreatedAt: time.Now(),
	}
	global.DB.Save(&user)
}

//Save，有主键记录就是更新，并且可以更新零值，否则就是创建

func update() {
	var user = models.UserModel{ID: 1}

	global.DB.Model(&user).
		Where("id=?", 1).
		Update("age", 21)

	fmt.Println(user)
}

func updateColumn() {
	global.DB = global.DB.Debug()
	var user = models.UserModel{ID: 1}
	global.DB.Model(&user).
		Where("id=?", 1).
		UpdateColumn("age", 22).
		UpdateColumn("name", "张三")
}

func updates() {
	var user = models.UserModel{ID: 2}

	//global.DB.Model(&user).
	//	Updates(models.UserModel{
	//		Name: "张三",
	//		Age:  22,
	//	})

	global.DB.Model(&user).
		Where("id = ?", 2).
		Updates(map[string]interface{}{
			"age": 23,
		})

	fmt.Println(user)
}

func main() {
	global.Connect()
	//save()
	//update()
	//updateColumn()
	updates()
}
