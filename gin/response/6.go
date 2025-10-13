package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("st", "static") //前面是别名，后面是实际路径
	r.StaticFile("abc", "static/download.html")
	r.Run(":8080")
}
