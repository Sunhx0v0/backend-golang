package main

import (
	"fmt"
	"server/mysqldb"
	"server/routes"
)

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
	// 调用输出化数据库的函数
	err := mysqldb.InitDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// 一次注册所有路由+授权跨域
	r := routes.SetupRoutes()
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("startup server failed,err: %v", err)
	}

	// r.Run()
	// r.Run(":8080")

}
