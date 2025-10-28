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
	err := DB.AutoMigrate(
		&models.UserModel{},
		&models.UserDetailModel{},
		&models.ClassModel{},
		&models.StudentModel{},
		&models.ArticleModel{},
		&models.TagModel{},
		&models.User1Model{},
		&models.Article1Model{},
		&models.User2ArticleModel{},
		&models.UserZdy{},
		&models.LogModel{})
	if err != nil {
		log.Fatalf("数据库迁移失败 %s", err)
	}
	fmt.Printf("数据库迁移成功")
}

func Connect() {
	dsn := "root:200088@tcp(127.0.0.1:3306)/gorm_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成实体外键
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}
