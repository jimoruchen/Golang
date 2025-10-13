package main

import "github.com/gin-gonic/gin"

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Index2(c *gin.Context) {
	c.JSON(200, response{
		Code: 0,
		Msg:  "成功",
		Data: map[string]any{},
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	//1.初始化
	r := gin.Default()
	//2.挂载路由
	r.GET("/index", Index2)
	//3.绑定端口，运行
	r.Run(":8080")
}
