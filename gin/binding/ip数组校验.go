package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type User struct {
			IPList []string `json:"ipList" binding:"dive,ip"`
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
