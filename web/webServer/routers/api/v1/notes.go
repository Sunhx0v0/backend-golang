package v1

import (
	"net/http"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

type Data struct {
	IsLogin bool          `json:"isLogin"` // 是否登录
	Notes   []models.Note `json:"notes"`   // 笔记，简要信息
}

func GetAllNotes(c *gin.Context) { //获取笔记（全部）
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

func GetSpecificNotes(c *gin.Context) { //获取特定笔记（搜索/标签）
	var data Data
	//判断是否登录，还要再加判断的函数
	data.IsLogin = false
	keyword := c.Param("keyword")
	data.Notes = models.GetSpBriefNtInfo(keyword)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}
