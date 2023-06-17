package v1

//用来写评论、点赞、收藏相关的处理函数
//以上动作都是通过在关系数据库添加记录完成的
import (
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

// 加载评论
func GetComments(c *gin.Context) {
	var comments []models.Comment
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	comments, success = models.GetCommentInfo(noteId)
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    comments,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "fail",
			"data":    comments,
		})
	}
}

// 发表评论
func PostComment(c *gin.Context) {
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	//获取前端传来的数据
	var newComment models.Comment
	//通过ShouldBind获取json数据
	if err := c.ShouldBind(&newComment); err == nil {
		success = models.NewComment(newComment, noteId)
		if success {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "评论成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "评论失败！",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
}

// 点赞某篇笔记
func LikeNote(c *gin.Context) {
	//数据库修改是否成功
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	var likeInfo models.LikeInfo
	//用shouldBind获取前端传来的json数据，只要json名相同就能读取
	if err := c.ShouldBind(&likeInfo); err == nil {
		//向数据库中插入点赞信息
		success = models.NewLike(likeInfo, noteId)
		if success {
			//将该笔记点赞数加一
			models.ChangeNoteLikes(noteId, 1)
			//将该笔记作者点赞数加一
			models.ChangeUserLikes(noteId, 1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "点赞成功！",
			})
		} else {
			//取消前面的点赞信息插入（好像可以省下来？）
			models.DeleteLike(likeInfo, noteId)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "点赞失败！",
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

// 取消点赞
func CancelLike(c *gin.Context) {
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	var likeInfo models.LikeInfo
	if err := c.ShouldBind(&likeInfo); err == nil {
		success = models.DeleteLike(likeInfo, noteId)
		if success {
			//将该笔记点赞数减一
			models.ChangeNoteLikes(noteId, -1)
			//将该笔记作者点赞数减一
			models.ChangeUserLikes(noteId, -1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "取消点赞成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "取消点赞失败！",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
}
