package main

import (
	"Golang/gin/response/res"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"code": 0,
		//	"msg":  "成功",
		//	"data": gin.H{},
		//})
		res.OkWithMsg(c, "登录成功")
	})

	r.GET("users", func(c *gin.Context) {
		res.OkWithData(c, map[string]any{
			"name": "zhangsan",
		})
	})

	r.POST("users", func(c *gin.Context) {
		res.FailWithMsg(c, "参数错误")
	})

	r.Run(":8080")
}
