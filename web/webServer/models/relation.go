package models

import (
	"fmt"
)

// 点赞信息
type LikeInfo struct {
	FvID        int `json:"fvId"`
	UserAct     int `json:"userAct"`
	FavorNoteID int `json:"favorNoteId"`
}

// 收藏信息
type CollectInfo struct {
	CltID         int `json:"cltId"`
	UserAct       int `json:"userAct"`
	CollectNoteID int `json:"collectNoteId"`
}

// @用户信息
type AtInfo struct {
	AtName     string `json:"atName" form:"atName"`
	AtLocation string `json:"atLocation" form:"atLocation"`
}

// 关注信息
type FollowRequest struct {
	FollowID string `json:"followID"` // 关注的人的id
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

// 插入收藏信息
func NewCollect(nclt CollectInfo, noteId int) bool {
	sqlstr := `INSERT INTO collectTable (userAct, collectNoteId) VALUES (?,?)`
	ret, err := db.Exec(sqlstr, nclt.UserAct, noteId)
	if err != nil {
		fmt.Printf("收藏信息insert failed, err:%v\n", err)
		return false
	}
	// 新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("收藏信息get lastinsert ID failed, err:%v\n", err)
		return false
	}
	fmt.Printf("收藏成功！收藏在数据库行数：%d\n", theID)
	return true
}

// 删除收藏信息
func DeleteCollect(dclt CollectInfo, noteId int) bool {
	sqlstr := "DELETE from collectTable WHERE userAct=? AND collectNoteId=?"
	ret, err := db.Exec(sqlstr, dclt.UserAct, noteId)
	if err != nil {
		fmt.Printf("收藏信息delete failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("收藏信息get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("收藏信息 delete success, affected rows:%d\n", n)
	return true
}

// 获取关注用户
func GetFollows() bool {
	return true
}

// 写入某篇笔记的@信息
func AddAtInfo(userId, noteId int, atinfos []AtInfo) bool {
	sqlstr := "INSERT INTO atTable (userAct, noteId, atUserName, atLocation) VALUES (?,?,?,?)"
	for _, atItem := range atinfos {
		_, err := db.Exec(sqlstr, userId, atItem.AtName, atItem.AtLocation)
		if err != nil {
			fmt.Printf("@信息insert failed, err:%v\n", err)
			return false
		}
	}
	return true
}

// 删除某篇笔记的@信息
func DeleteAtInfo(noteId int) bool {
	sqlstr := "DELETE from atTable WHERE noteId=?"
	ret, err := db.Exec(sqlstr, noteId)
	if err != nil {
		fmt.Printf("@信息delete failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("@信息get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("@信息 delete success, affected rows:%d\n", n)
	return true
}

// 增加关注信息
func AddFollowInfo(useraccount int, account int) bool {
	sqlstr := "insert into followTable (userAct,followAct) values (?,?)"
	ret, err := db.Exec(sqlstr, useraccount, account)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("收藏信息get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("收藏信息 delete success, affected rows:%d\n", n)
	return true
}

// 删除关注信息
func DelFollowInfo(useraccount int, account int) bool {
	sqlstr := "delete from followTable where userAct=? and followAct=?"
	ret, err := db.Exec(sqlstr, useraccount, account)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("收藏信息get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("收藏信息 delete success, affected rows:%d\n", n)
	return true
}

// 修改关注数
func ChangeUserFollows(userId, option int) {
	var sqlstr string
	addnum := `UPDATE userInfo set followNum =followNum+1 WHERE userAccount = ?`
	reducenum := `UPDATE userInfo set followNum =followNum-1 WHERE userACCount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("关注数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("关注数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("关注数修改编号：%d\n", n)
}

// 修改粉丝数
func ChangeUserFans(userId, option int) {
	var sqlstr string
	addnum := `UPDATE userInfo set fansNum =fansNum+1 WHERE userAccount = ?`
	reducenum := `UPDATE userInfo set fansNum =fansNum-1 WHERE userAccount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("粉丝数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("粉丝数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("粉丝数修改编号：%d\n", n)
}
