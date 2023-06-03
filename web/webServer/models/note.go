package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// type Article struct {
// 	Model

// 	State         int    `json:"state"`
// 	TagId         int    `json:"tag_id"`
// 	Title         string `json:"title"`
// 	Desc          string `json:"desc"`
// 	Content       string `json:"Content"`
// 	CoverImageUrl string `json:"cover_image_url"`
// 	CreatedBy     string `json:"created_by"`
// 	Tag           Tag    `json:"tag"`
// }

type noteInfo struct {
	NoteId    int    `JSON:"NoteId"`
	NoteTitle string `JSON:"NoteTitle"`
	NoteCover string `JSON:"NoteCover"`
}

var notes []noteInfo

// 查询多条数据示例
func queryNoteDemo() []noteInfo {
	sqlStr := "select noteId, title, cover from noteInfo where noteId > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt noteInfo
		err := rows.Scan(&nt.NoteId, &nt.NoteTitle, &nt.NoteCover)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		notes = append(notes, nt)
		//fmt.Printf("id:%d name:%s age:%s\n", nt.noteId, nt.noteTitle, nt.noteCover)
	}
	return notes
}

// func getNotes(PageNum int, PageSize int, maps interface{}) (article []Article) {
// 	db.Where(maps).Offset(PageNum).Limit(PageSize).Find(&article)
// 	return
// }

// func AddArticle(data map[string]interface{}) bool {
// 	db.Create(&Article{
// 		TagId:     data["tag_id"].(int),
// 		Title:     data["title"].(string),
// 		Desc:      data["desc"].(string),
// 		Content:   data["content"].(string),
// 		CreatedBy: data["created_by"].(string),
// 		State:     data["state"].(int),
// 	})

// 	return true
// }

// func EditArticle(id int, maps interface{}) bool {
// 	db.Model(&Article{}).Where("id = ?", id).Update(maps)
// 	return true
// }

// func DeleteArticle(id int) bool {
// 	db.Where("id = ?", id).Delete(&Article{})
// 	return true
// }

// func ExistArticleByID(id int) bool {
// 	var article Article
// 	db.Select("id").Where("id = ?", id).First(&article)
// 	if article.ID > 0 {
// 		return true
// 	}
// 	return false
// }

// func GetArticleTotal(maps interface{}) (count int) {
// 	db.Model(&Article{}).Where(maps).Count(&count)
// 	return
// }
