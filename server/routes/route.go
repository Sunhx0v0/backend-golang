package routes

import (
	"net/http"
	"server/routes/admin"
	"server/routes/web"

	"github.com/gin-gonic/gin"
)

// 处理跨域请求,将跨域请求函数作为中间件处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")                                                                                         // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")                              //header的类型
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                                                          //允许请求方法
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type") //返回数据格式
		c.Header("Access-Control-Allow-Credentials", "true")                                                                                 //设置为true，允许ajax异步请求带cookie信息

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	// 全局使用中间件（授权跨域请求）
	r.Use(Cors())
	admin.SetupAdminRoutes(r)
	web.SetupWebRoutes(r)
	// 一次加载所有html（不一定用得上）
	r.LoadHTMLGlob("./templates/**/*")
	return r
}
