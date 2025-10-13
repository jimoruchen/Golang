package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"fmt"
)

func main() {
	global.Connect()

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
	userList := []models.UserModel{
		{Name: "Alex", Age: 18},
		{Name: "Tom", Age: 20},
		{Name: "Invalid", Age: -5}, // 假设这个会触发数据库约束失败
	}
	for _, user := range userList {
		err := global.DB.Create(&user).Error
		if err != nil {
			fmt.Printf("创建用户 %s 失败：%v", user.Name, err)
			continue
		}
		fmt.Printf(" 创建用户 %s 成功", user.Name)
	}

}
