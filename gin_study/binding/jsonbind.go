package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("json", func(c *gin.Context) {
		type User struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
