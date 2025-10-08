package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		name := c.PostForm("name")
		age, ok := c.GetPostForm("age") //判断有没有传
		fmt.Println(name)
		fmt.Println(age, ok)
	})
	r.Run(":8080")
}
