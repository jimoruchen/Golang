package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"encoding/json"
	"fmt"
)

func test() {
	//global.DB.Create(&models.LogModel{
	//	Title: "用户注册",
	//	Level: models.InfoLevel,
	//})

	var log models.LogModel
	global.DB.Take(&log)
	byteData, _ := json.Marshal(log)
	fmt.Println(string(byteData))
}

func main() {
	global.Connect()
	//global.Migrate()
	test()
}
