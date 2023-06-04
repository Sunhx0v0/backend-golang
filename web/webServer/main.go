// 主页面
package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	"webServer/models"

	"webServer/routers"

	_ "github.com/go-sql-driver/mysql"
)

func hostFunc(c *gin.Context) {
	// gin.H 是map[string]interface{}的缩写
	c.HTML(http.StatusOK, "host.html", gin.H{
		"title": "标题",
		"nt":    models.Notes,
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang http!")
}

func main() {
	err := models.InitDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	models.QueryNoteDemo()

	router := gin.Default()

	//加载HTML文件
	router.LoadHTMLGlob("templates/host.html")
	router.GET("/", hostFunc)
	router.Run(":8080")
}
