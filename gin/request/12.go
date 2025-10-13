package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _raw(c *gin.Context, obj any) (err error) {
	byteData, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(byteData, &obj)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(obj)
	}
	return nil
}

func updateUser(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var user User
	err := _raw(c, &user)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, Response{0, user, "成功"})
}

func main() {
	r := gin.Default()
	r.POST("/", updateUser)
	r.Run(":8080")
}
