package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	dsn := "root:200088@tcp(127.0.0.1:3306)/gorm_db?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(db)

	var userList []User
	db.Find(&userList)
	fmt.Println(userList)

	db.Where("id < ?", 2).Find(&userList)
	fmt.Println(userList)

	db.Where("name = ?", "李四").Find(&userList)
	fmt.Println(userList)
}
