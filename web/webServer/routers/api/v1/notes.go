package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webServer/models"
)

type Data struct {
	IsLogin bool          `json:"isLogin"` // 是否登录
	Notes   []models.Note `json:"notes"`   // 笔记，简要信息
}

func GetAllNotes(c *gin.Context) {
	var data Data
	//判断是否登录，还要再加判断的函数
	data.IsLogin = false
	data.Notes = models.GetBriefNtInfo()
	// gin.H 是map[string]interface{}的缩写
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}
