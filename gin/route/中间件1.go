package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home...")
	c.String(200, "Home")
}

func M1(c *gin.Context) {
	fmt.Println("M1请求部分")
	c.Next()
	fmt.Println("M1响应部分")
}

func M2(c *gin.Context) {
	fmt.Println("M2请求部分")
	c.Next()
	fmt.Println("M2响应部分")
}

func main() {
	r := gin.Default()
	r.GET("/", M1, M2, Home)

	r.Run(":8080")
}
