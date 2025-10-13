package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 身份验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供认证信息"})
			c.Abort() // 终止请求处理链
			return
		}
		// 这里可以添加更复杂的验证逻辑，比如 JWT 解析
		if token != "my-secret-token" {
			c.JSON(http.StatusForbidden, gin.H{"error": "认证失败"})
			c.Abort()
			return
		}

		// 认证通过，继续执行
		c.Next()
	}
}

// MetricsMiddleware 另一个中间件，用于记录特定路由的访问
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录访问次数等指标
		c.Next()
	}
}

func main() {
	r := gin.New()

	// --- 局部中间件应用到单个路由 ---
	r.GET("/public", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "公开接口，无需认证"})
	})

	// 这个路由需要认证
	r.GET("/private", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "私有接口，认证成功"})
	})

	// --- 局部中间件应用到路由组 ---
	authorized := r.Group("/admin", AuthMiddleware(), MetricsMiddleware())
	{
		authorized.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "管理员仪表盘"})
		})

		authorized.POST("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "创建用户"})
		})
	}

	// 注意：/public 路由不会经过 AuthMiddleware 或 MetricsMiddleware

	r.Run(":8080")
}
