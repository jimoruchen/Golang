package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	r := route.Group("/api")
	r.POST("/users", func(c *gin.Context) {
		url := c.Request.URL
		fmt.Println(url, c.Request.Method)
	})

	route.Run(":8080")
}
