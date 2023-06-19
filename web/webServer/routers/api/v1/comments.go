package v1

import (
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

// 笔记页面加载评论
func GetComments(c *gin.Context) {
	var comments []models.Comment
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	comments, success = models.GetCommentInfo(noteId, 0)
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
			//将该笔记点赞数加一
			models.ChangeNoteComments(noteId, 1)
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

// 删除评论
func CancleComment(c *gin.Context) {
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	var comment models.Comment
	if err := c.ShouldBind(&comment); err == nil {
		success = models.DeleteComment(int(comment.CommentID))
		if success {
			//将该笔记点赞数加一
			models.ChangeNoteComments(noteId, -1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "评论删除成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "取消评论失败！",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
}

// 消息列表加载评论
func MsgGetComments(c *gin.Context) {
	var comments []models.Comment
	var success bool
	userId, _ := strconv.Atoi(c.Param("userId"))
	comments, success = models.GetCommentInfo(userId, 1)
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

// 修改某条评论状态
func ChangeCommentState(c *gin.Context) {
	var success bool
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	success = models.SetCommentState(commentId)
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "评论状态修改失败",
		})
	}
}
