package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userRegister(r *gin.Engine) {
	r.GET("/user/index", indexHandler)
}

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		// c.JSON：返回 JSON 格式的数据
		"message": "Hello, this is user-index",
	})
	// c.HTML(http.StatusOK, "server/templates/users_test/index.html", gin.H{
	// 	"title": "users/index",
	// })
}
