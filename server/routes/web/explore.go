package web

import (
	"net/http"
	"server/mysqldb"
	"strconv"

	"github.com/gin-gonic/gin"
)

func exploreRegister(r *gin.Engine) {
	r.GET("/explore", exploreHandler)
	r.GET("/explore/:resource", NoteDetailHandler)
}

func exploreHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		// c.JSON：返回 JSON 格式的数据
		"message": "Hello, this is user-index",
	})
	// c.HTML(http.StatusOK, "server/templates/users_test/index.html", gin.H{
	// 	"title": "users/index",
	// })
}

func NoteDetailHandler(c *gin.Context) {
	noteid, _ := strconv.Atoi(c.Param("resource"))
	data := mysqldb.SpecificNote(noteid)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
	// c.HTML(http.StatusOK, "server/templates/users_test/index.html", gin.H{
	// 	"title": "users/index",
	// })
}
