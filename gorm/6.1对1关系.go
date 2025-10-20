package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
)

func oneToOne() {
	//创建用户，连带着创建用户详情
	//err := global.DB.Create(&models.UserModel{
	//	Name: "张三",
	//	Age:  18,
	//	UserDetailModel: &models.UserDetailModel{
	//		Email: "111@qq.com",
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//创建用户详情，关联用户
	//err := global.DB.Create(&models.UserDetailModel{
	//	Email: "222@qq.com",
	//	//UserID: 13,
	//	UserModel: &models.UserModel{ID: 13},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//通过用户详情查用户		正查
	//var id = 12
	//var detail models.UserDetailModel
	//global.DB.Preload("UserModel").Take(&detail, "user_id = ?", id)
	//fmt.Println(detail.Email, detail.UserModel.Name)

	//反查
	//var id = 12
	//var user models.UserModel
	//global.DB.Preload("UserDetailModel").Take(&user, id)
	//fmt.Println(user, user.UserDetailModel.Email)

	//级联删除
	//var user models.UserModel
	//global.DB.Find(&user, 12)
	//fmt.Println(user)
	//global.DB.Select("UserDetailModel").Delete(&user)

	//set null
	var user models.UserModel
	global.DB.Find(&user, 13)
	global.DB.Model(&user).Association("UserDetailModel").Clear()
	global.DB.Delete(&user)
}

func main() {
	global.Connect()
	//global.Migrate()
	oneToOne()
}
