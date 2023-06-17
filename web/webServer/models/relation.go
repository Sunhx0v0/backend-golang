package models

import (
	"fmt"
)

// 评论信息
type Comment struct {
	CommentatorID   int64  `json:"commentatorId"`   // 评论者账号
	CommentatorName string `json:"commentatorName"` // 评论者
	CommentID       int64  `json:"commentId"`       // 评论编号
	CommentTime     string `json:"commentTime"`     // 评论时间
	Content         string `json:"content"`         // 评论内容
	Portrait        string `json:"portrait"`        // 用户头像
}

// 点赞信息
type LikeInfo struct {
	FvID        int `json:"fvId"`
	UserAct     int `json:"userAct"`
	FavorNoteID int `json:"favorNoteId"`
}

// 根据笔记编号获取作者账号
func NoteToUser(noteId int) int {
	var userId int
	sltsql := "SELECT creatorAccount FROM noteInfo WHERE noteId=?"
	err := db.QueryRow(sltsql, noteId).Scan(&userId)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return -1
	} else {
		return userId
	}
}

// 获取评论
func GetCommentInfo(noteId int) (comments []Comment, ok bool) {
	ok = true
	sqlstr := `SELECT c.commentId, c.commentatorId, c.content, c.commentTime, u.userName, u.portrait
	FROM commentInfo c, userInfo u
	WHERE c.noteID = ? AND c.commentatorId = u.userAccount`
	rows, err := db.Query(sqlstr, noteId)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		ok = false
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var cmt Comment
		err := rows.Scan(&cmt.CommentID, &cmt.CommentatorID, &cmt.Content, &cmt.CommentTime, &cmt.CommentatorName, &cmt.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			ok = false
			return
		}
		comments = append(comments, cmt)
	}
	return comments, ok
}

// 插入评论信息
func NewComment(nc Comment, noteId int) bool {
	sqlstr := `INSERT INTO commentInfo (noteID, commentatorId, content, commentTime) VALUES (?,?,?,?)`
	ret, err := db.Exec(sqlstr, noteId, nc.CommentatorID, nc.Content, nc.CommentTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	// 新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return false
	}
	fmt.Printf("评论成功！评论在数据库行数：%d\n", theID)
	return true
}

// 插入点赞信息
func NewLike(nl LikeInfo, noteId int) bool {
	sqlstr := `INSERT INTO favorTable (userAct, favorNoteId) VALUES (?,?)`
	ret, err := db.Exec(sqlstr, nl.UserAct, noteId)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	// 新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return false
	}
	fmt.Printf("点赞成功！评论在数据库行数：%d\n", theID)
	return true
}

// 删除点赞信息
func DeleteLike(nl LikeInfo, noteId int) bool {
	// userId := NoteToUser(noteId)
	sqlstr := "DELETE from favorTable WHERE userAct=? AND favorNoteId=?"
	ret, err := db.Exec(sqlstr, nl.UserAct, noteId)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("点赞信息 delete success, affected rows:%d\n", n)
	return true
}
