package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("header", func(c *gin.Context) {
		type User struct {
			Name        string `header:"name"`
			Id          int    `header:"id"`
			UserAgent   string `header:"user-agent"`
			ContentType string `header:"content-type"`
		}
		var user User
		err := c.ShouldBindHeader(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
