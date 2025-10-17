package global

import (
	"Golang/gorm/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate() {
	err := DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		log.Fatalf("数据库迁移失败 %s", err)
	}
	fmt.Printf("数据库迁移成功")
}

func Connect() {
	dst := "root:200088@tcp(127.0.0.1:3303)/gorm_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}
