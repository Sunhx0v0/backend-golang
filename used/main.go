package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func hostFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from API!",
		"info":    "终于tnnd成功了!!!",
	})
}

// 用于处理Web应用程序中的跨源资源共享（CORS）请求的中间件，不用管，要看也行。模板写法来着，至于为什么，我也不是很懂
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware()) // 一定要导入，我也不懂为什么
	// 提供静态资源
	r.StaticFS("/static", http.Dir("./myproject-vue/dist")) // 获取vue的静态界面（没什么用，或许可用于初始化界面）

	// 处理API请求
	r.GET("/api/hello", hostFunc) // 动态处理vue端的API请求

	// 启动服务
	r.Run(":8083") //  如果只启动后端而不启动前端，那么后端无法获取前段的静态界面
}
