package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func page1(c *gin.Context) {
	c.HTML(http.StatusOK, "template1.html", gin.H{
		"title": "标题",
	})
}

func main() {
	router := gin.Default()
	//加载.tmpl文件
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	//加载HTML文件
	router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/template1", page1)
	router.Run(":8080")
}
