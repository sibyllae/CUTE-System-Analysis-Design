package main

import (
	"cute/api"
	"cute/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化 Gin 引擎
	r := gin.Default()

	// 配置 CORS 中间件，允许所有请求
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 初始化数据库

	database.Init()

	api.RegisterPingRoutes(r)
	api.RegisterUserRoutes(r)
	api.RegisterGameRoutes(r)
	api.RegisterGameTypeRoutes(r)
	api.RegisterLoginRecordRoutes(r)
	api.RegisterOrderRoutes(r)

	// 啟動 HTTP 服務，監聽 8080 端口
	r.Run(":8080")
}
