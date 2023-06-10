package v1

import (
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Infos models.UserInfo `json:"userInfo"` // 用户信息，只有一条，不用数组
	Notes []models.Note   `json:"notes"`    // 笔记，简要信息
}

func GetUserInfo(c *gin.Context) { //显示用户界面全部信息
	var info UserInfo
	//判断是否登录，还要再加判断的函数
	userID, _ := strconv.Atoi(c.Param("userID"))
	info.Infos = models.UserInfoDB(userID) // 通过用户ID去数据库获取信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    info.Infos,
	})
}
