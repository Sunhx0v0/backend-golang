package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// type noteInfo struct {
// 	NoteId    int    `JSON:"NoteId"`
// 	NoteTitle string `JSON:"NoteTitle"`
// 	NoteCover string `JSON:"NoteCover"`
// }

// var Notes []noteInfo

type Note struct {
	Cover     string `json:"cover"`
	CreatorID int    `json:"creatorId"` // 作者编号
	LikedNum  int    `json:"likedNum"`  // 点赞数
	NoteID    int    `json:"noteId"`    // 笔记编号
	Portrait  string `json:"portrait"`  // 头像
	Title     string `json:"title"`
}

// 获取笔记的封面标题等简要信息
func GetBriefNtInfo() (notes []Note) {
	sqlStr := `select n.noteId, n.title, n.cover,n.creatorAccount,n.likeNum, u.portrait
	from noteInfo n,userInfo u
	where n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Print(nt.NoteID)
		notes = append(notes, nt)
	}
	return
}

// 获取特定内容的笔记
func GetSpBriefNtInfo(keyword string) (notes []Note) {
	sqlStr := `select n.noteId, n.title, n.cover,n.creatorAccount,n.likeNum, u.portrait
	from noteInfo n,userInfo u
	where n.creatorAccount = u.userAccount and (n.tag = ? OR n.title like CONCAT('%',#{keyword},'%'))`
	rows, err := db.Query(sqlStr, keyword, keyword)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Print(nt.NoteID)
		notes = append(notes, nt)
	}
	return
}
