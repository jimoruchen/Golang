package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("users/:id/:name", func(c *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			Id   int    `uri:"id"`
		}
		var user User
		err := c.ShouldBindUri(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
