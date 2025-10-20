package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"fmt"
)

func oneToMany() {
	err := global.DB.Create(&models.ClassModel{
		Name: "class1",
		StudentList: []models.StudentModel{
			{Name: "student1"},
			{Name: "student2"},
		},
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	global.Connect()
	//global.Migrate()
	oneToMany()
}
