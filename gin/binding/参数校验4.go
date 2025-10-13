package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type User struct {
			FileName string `json:"filename" binding:"endswith=.jpg"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
