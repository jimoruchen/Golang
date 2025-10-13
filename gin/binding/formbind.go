package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("form", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Id   int    `form:"id"`
		}
		var user User
		err := c.ShouldBind(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
