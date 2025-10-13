package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 一个简单的日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Printf("请求开始: %s %s\n", c.Request.Method, c.Request.URL.Path)

		// 在中间件中设置变量，可以在后续的处理器中获取
		c.Set("example", "全局中间件设置的值")

		// 执行下一个中间件或最终的处理器
		c.Next()

		// 请求处理完成后执行
		latency := time.Since(t)
		fmt.Printf("请求结束，耗时: %v\n", latency)
	}
}

func Logger1(c *gin.Context) {
	t := time.Now()
	fmt.Printf("请求开始: %s %s\n", c.Request.Method, c.Request.URL.Path)

	// 在中间件中设置变量，可以在后续的处理器中获取
	c.Set("example", "全局中间件设置的值")

	// 执行下一个中间件或最终的处理器
	c.Next()

	// 请求处理完成后执行
	latency := time.Since(t)
	fmt.Printf("请求结束，耗时: %v\n", latency)
}

func main() {
	r := gin.Default() // Default() 默认包含了 Logger 和 Recovery 中间件

	// 注册全局中间件
	//r.Use(Logger())
	r.Use(Logger1) //不能加括号

	// 定义一些路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/user", func(c *gin.Context) {
		// 可以获取全局中间件中设置的值
		value, exists := c.Get("example")
		if exists {
			fmt.Println("从中间件获取的值:", value)
		}
		c.JSON(200, gin.H{"user": "admin"})
	})

	r.Run(":8080") // 监听并在 0.0.0.0:8080 启动服务
}
