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
func AddLikeInfo(useraccount int, noteid int) {
	sqlstr := "insert into favorTable (userAct,favorNoteId) values (?,?)"
	_, err := db.Exec(sqlstr, useraccount, noteid)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
}

// 增加关注信息
func AddFollowInfo(useraccount int, account int) {
	sqlstr := "insert into followTable (userAct,followAct) values (?,?)"
	_, err := db.Exec(sqlstr, useraccount, account)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
}

// 增加收藏信息
func AddCollectInfo(useraccount int, noteid int) {
	sqlstr := "insert into collectTable (userAct,collectNoteId) values (?,?)"
	_, err := db.Exec(sqlstr, useraccount, noteid)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
}

// ——————————————————————————————————————————————————————————————————————————————————
// 删除点赞信息
func DelLikeInfo(useraccount int, noteid int) {
	sqlstr := "delete from favorTable where userAct=? and favorNoteId=?"
	_, err := db.Exec(sqlstr, useraccount, noteid)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
}

// 删除关注信息
func DelFollowInfo(useraccount int, account int) {
	sqlstr := "delete from followTable where userAct=? and followAct=?"
	_, err := db.Exec(sqlstr, useraccount, account)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
}

// 删除收藏信息
func DelCollectInfo(useraccount int, noteid int) {
	sqlstr := "delete from collectTable where userAct=? and collectNoteId=?"
	_, err := db.Exec(sqlstr, useraccount, noteid)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
}
