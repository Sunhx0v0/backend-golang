package mysqldb

import (
	"fmt"
	"server/models"
)

type LikeInfo models.LikeTable
type FollowInfo models.FollowTable
type CollectInfo models.CollectTable

// //按用户查询【三连】信息
// func TripleByUserAccount(useraccount int) ([]int, []Note, []Note) {

// }
// ——————————————————————————————————————————————————————————————————————————————————
// 增加点赞信息
func AddLikeInfo(useraccount int, noteid int) bool {
	sqlstr := "insert into favorTable (userAct,favorNoteId) values (?,?)"
	ret, err := db.Exec(sqlstr, useraccount, noteid)
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

// 增加收藏信息
func AddCollectInfo(useraccount int, noteid int) bool {
	sqlstr := "insert into collectTable (userAct,collectNoteId) values (?,?)"
	ret, err := db.Exec(sqlstr, useraccount, noteid)
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

// ——————————————————————————————————————————————————————————————————————————————————
// 删除点赞信息
func DelLikeInfo(useraccount int, noteid int) bool {
	sqlstr := "delete from favorTable where userAct=? and favorNoteId=?"
	ret, err := db.Exec(sqlstr, useraccount, noteid)
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

// 删除收藏信息
func DelCollectInfo(useraccount int, noteid int) bool {
	sqlstr := "delete from collectTable where userAct=? and collectNoteId=?"
	ret, err := db.Exec(sqlstr, useraccount, noteid)
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
