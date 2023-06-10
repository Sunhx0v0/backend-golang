package v1

import (
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Infos    []models.UserInfo `json:"userInfo"` // 用户信息，只有一条，不用数组
	Notes    []models.Notes    `json:"notes"`    // 笔记，简要信息
	Collects []models.Collects `json:"collects"`
	Likes    []models.Likes    `json:"likes"`
	IsHost   bool              `json:"isHost"` //是否页面主人
}

func GetUserInfo(c *gin.Context) { //显示用户界面全部信息
	var info UserInfo
	userID, _ := strconv.Atoi(c.Param("userID"))
	info.Infos = models.UserInfoDB(userID) // 通过用户ID去数据库获取信息
	info.Notes = models.NoteInfoDB(userID)
	info.Collects = models.CollectInfoDB(userID)
	info.Likes = models.LikeInfoDB(userID)
	info.IsHost = true

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    info,
	})
}
