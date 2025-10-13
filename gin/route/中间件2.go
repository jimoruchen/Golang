package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home1(c *gin.Context) {
	fmt.Println("Home...")
	c.String(200, "Home")
}

func M3(c *gin.Context) {
	fmt.Println("M3请求部分")
	c.Abort()
	fmt.Println("M3响应部分")
}

func main() {
	r := gin.Default()
	r.GET("/", M3, Home1)

	r.Run(":8080")
}
