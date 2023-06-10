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

func UploadNote(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))

	form, err := c.MultipartForm()
	files := form.File["files[]"]
	if err != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "上传失败!",
		})
		return
	} else {
		//新声明新笔记的结构体
		var newNote models.DetailNote
		newNote.Title = c.Query("title")
		newNote.Body = c.Query("body")
		newNote.CreateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", c.Query("createtime"), time.Local)
		newNote.UpdateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", c.Query("createtime"), time.Local)
		newNote.Tag = c.Query("tag")
		newNote.Location = c.Query("location")
		newNote.AtUserID = com.StrTo(c.Query("atuserid")).MustInt()
		newNote.LikedNum = com.StrTo(c.Query("likenum")).MustInt()

		picNum := 0
		newNote.CreatorID = userId
		newNote.Picnum = picNum
		ntID, success := models.NewNoteInfo(newNote)
		if !success {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "上传失败!",
			})
			return
		}
		for index, file := range files {
			var pc models.Pictures

			log.Println(file.Filename)
			dst := fmt.Sprintf("../images/%s_%d", file.Filename, index)

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
		// fileExt := strings.ToLower(path.Ext(f.Filename))
		// if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		// 	c.JSON(200, gin.H{
		// 		"code":    400,
		// 		"message": "上传失败!只允许png,jpg,gif,jpeg文件",
		// 	})
		// 	return
		// }
		// fileName := tools.Md5(fmt.Sprintf("%s%s", f.Filename, time.Now().String()))
		// fildDir := fmt.Sprintf("%s%d%s/", config.Upload, time.Now().Year(), time.Now().Month().String())
		// isExist, _ := tools.IsFileExist(fildDir)
		// if !isExist {
		// 	os.Mkdir(fildDir, os.ModePerm)
		// }
		// filepath := fmt.Sprintf("%s%s%s", fildDir, fileName, fileExt)
		// c.SaveUploadedFile(f, filepath)
		// c.JSON(200, gin.H{
		// 	"code":    200,
		// 	"message": "上传成功!",
		// 	"result": gin.H{
		// 		"path": filepath,
		// 	},
		// })
	}
}
