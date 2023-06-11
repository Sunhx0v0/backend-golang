package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"webServer/models"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type Data struct {
	IsLogin bool          `json:"isLogin"` // 是否登录
	Notes   []models.Note `json:"notes"`   // 笔记，简要信息
}

// 获取笔记（全部）
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

// 获取特定笔记（搜索/标签）
func GetSpecificNotes(c *gin.Context) {
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

// 上传笔记
func UploadNote(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "读取失败!",
		})
		return
	} else {
		files := form.File["files"]
		//新声明新笔记的结构体
		var newNote models.DetailNote
		newNote.Title = c.PostForm("title")
		newNote.Body = c.PostForm("body")
		newNote.CreateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", c.PostForm("createtime"), time.Local)
		newNote.UpdateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", c.PostForm("createtime"), time.Local)
		newNote.Tag = c.PostForm("tag")
		newNote.Location = c.PostForm("location")
		newNote.AtUserID = com.StrTo(c.PostForm("atuserid")).MustInt()
		// newNote.LikedNum = com.StrTo(c.PostForm("likenum")).MustInt()

		// picNum := 0
		newNote.CreatorID = userId
		// newNote.Picnum = picNum
		ntID, success := models.NewNoteInfo(newNote)
		if !success {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "上传失败!",
			})
			return
		}
		for index, file := range files {
			var pc models.Pictures

			log.Println(file.Filename)
			dst := fmt.Sprintf("images/%d_%s", index, file.Filename)

			pc.NoteId = ntID
			pc.Picurl = dst
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
			//将路径等信息更新到数据库
			models.NewPicInfo(pc)
			// picNum++
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	}
}

// 删除笔记
func DeleteNote(c *gin.Context) {

}
