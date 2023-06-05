package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webServer/models"
)

func HostFunc(c *gin.Context) {
	models.QueryNoteDemo()
	// gin.H 是map[string]interface{}的缩写
	c.HTML(http.StatusOK, "host.html", gin.H{
		"title": "标题",
		"nt":    models.Notes,
	})
}
