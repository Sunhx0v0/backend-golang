package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
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
			"message": fmt.Sprintf("读取失败! err:%s", err),
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
				"message": "表单内容获取失败!",
			})
			return
		}

		fileType := map[string]bool{
			".png":  true,
			".jpg":  true,
			".jpeg": true,
			".gif":  true,
		}
		for _, file := range files {
			extName := path.Ext(file.Filename)
			_, b := fileType[extName]

			if !b {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"message": "上传文件类型不合法！",
				})
				return
			}
			var pc models.Pictures
			timeStamp := time.Now().Unix()

			log.Println(file.Filename)
			dst := fmt.Sprintf("images/%d_%d_%s_%s", userId, ntID, strconv.Itoa(int(timeStamp)), file.Filename)

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

	userId, _ := strconv.Atoi(c.Param("userId"))
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	files, err := Getfile(userId, noteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "获取对应信息失败",
		})
		return
	}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "删除失败",
			})
			return
		}
	}
	if models.DeletePic(noteId) && models.DeleteNote(noteId) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "删除成功！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "数据库删除成功！",
		})
	}
}

func Getfile(userid, noteid int) ([]string, error) {
	var files []string
	f, err := os.Open("images")
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}
	filter := fmt.Sprintf("%d_%d", userid, noteid)
	for _, file := range fileInfo {
		if strings.Contains(file.Name(), filter) {
			files = append(files, fmt.Sprintf("%s%s", "images/", file.Name()))
		}
	}
	return files, nil
}

// func walkFunc(path string, info os.FileInfo, err error) error {
//     if err != nil {
//         // 错误处理
//         return err
//     }
//     if !info.IsDir() && strings.Contains(path, ".txt") {
//         // 处理符合条件的文件
//         fmt.Println(path)
//     }
//     return nil
// }
