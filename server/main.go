package main

import (
	"fmt"
	"server/mysqldb"
	"server/routes"
)

func main() {
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
