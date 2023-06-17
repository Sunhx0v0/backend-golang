package web

import (
	"net/http"
	"server/mysqldb"
	"strconv"

	"github.com/gin-gonic/gin"
)

func userRegister(r *gin.Engine) {
	r.GET("/user/index", indexHandler)
	r.POST("/:userId/PersonalView/follow", followHandler)
	r.DELETE("/:userId/PersonalView/follow", cancelFollowHandler)
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

type Request struct {
	FollowID string `json:"followID"` // 关注的人的id
}

func followHandler(c *gin.Context) {
	var success bool
	userId, _ := strconv.Atoi(c.Param("userId"))
	var account Request
	//用shouldBind获取前端传来的json数据，只要json名相同就能读取
	if err := c.ShouldBind(&account); err == nil {
		id, _ := strconv.Atoi(account.FollowID)
		//向数据库中插入关注信息
		success = mysqldb.AddFollowInfo(userId, id)
		if success {
			//将用户关注数加一
			mysqldb.ChangeUserFollows(userId, 1)
			//将被关注用户粉丝数加一
			mysqldb.ChangeUserFans(userId, 1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "关注成功！",
			})
		}
	} else {
		//json数据获取失败
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  40,
			"error": err.Error(),
		})
	}
}

// 取消收藏
func cancelFollowHandler(c *gin.Context) {
	var success bool
	userId, _ := strconv.Atoi(c.Param("userId"))
	var account int
	//用shouldBind获取前端传来的json数据，只要json名相同就能读取
	if err := c.ShouldBind(&account); err == nil {
		//向数据库中插入关注信息
		success = mysqldb.DelFollowInfo(userId, account)
		if success {
			//将用户关注数加一
			mysqldb.ChangeUserFollows(userId, -1)
			//将被关注用户粉丝数加一
			mysqldb.ChangeUserFans(userId, -1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "关注成功！",
			})
		}
	} else {
		//json数据获取失败
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
}
