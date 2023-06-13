package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

type UsersInfo struct {
	Infos    models.UserInfo `json:"userInfo"` // 用户信息，只有一条，不用数组
	Notes    []models.Notes  `json:"notes"`    // 笔记，简要信息
	Collects []models.Notes  `json:"collects"`
	Likes    []models.Notes  `json:"likes"`
	IsHost   bool            `json:"isHost"` //是否页面主人
}

func GetUserInfo(c *gin.Context) { //显示用户界面全部信息
	var info UsersInfo
	userID, _ := strconv.Atoi(c.Param("userId"))
	fmt.Println("用户ID:", userID)
	// 通过用户ID去数据库获取信息
	info.Infos = models.UserInfoDB(userID)
	// 获取某用户发布的笔记
	info.Notes = models.NoteInfoDB(userID)
	// 获取某用户收藏的笔记
	info.Collects = models.CollectInfoDB(userID)
	// 获取某用户点赞的笔记
	info.Likes = models.LikeInfoDB(userID)
	info.IsHost = true
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    info,
	})
}
