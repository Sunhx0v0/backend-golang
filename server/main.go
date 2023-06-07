package main

import (
	"fmt"
	"net/http"
	"server/routes"

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

func main() {
	fmt.Println("这是第一个代码")
	fmt.Println("在【终端】执行了")
	fmt.Println("go mod init server")
	fmt.Println("go build")
	// // 创建一个默认的路由引擎
	// r := gin.Default()
	// // 加载模板
	// r.LoadHTMLFiles("templates/posts_test/index.html", "templates/users_test/index.html", "templates/test/tnt_web.html")
	// // 配置路由
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		// c.JSON：返回 JSON 格式的数据
	// 		"message": "Hello world!",
	// 	})
	// })
	// r.GET("/posts/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "posts_test/index.html", gin.H{
	// 		"title": "posts/index",
	// 	})
	// })
	// r.GET("/users/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "users_test/index.html", gin.H{
	// 		"title": "users/index",
	// 	})
	// })
	// r.GET("/tnt/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "test/tnt_web.html", gin.H{
	// 		"title": "tnt/index",
	// 	})
	// })
	// // 接收页面form参数
	// r.POST("/tnt/index", func(c *gin.Context) {
	// 	tnt_url := c.PostForm("url_list")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": tnt_url,
	// 	})
	// })

	// 一次注册所有路由
	r := routes.SetupRoutes()
	// 全局使用中间件（授权跨域请求）
	r.Use(Cors())
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("startup server failed,err: %v", err)
	}

	// r.Run()
	// r.Run(":8080")

}
