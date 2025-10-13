package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("gin/response/static/*")  	//加载所有
	r.LoadHTMLFiles("gin/response/static/index.html") //加载单个
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", map[string]any{
			"title": "即墨如尘",
		})
	})
	r.Run(":8080")
}
