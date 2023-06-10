package models

import (
	"fmt"
)

type Pictures struct {
	PicID  string `json:"pic"`
	NoteId int    `json:"noteid"`
	Picurl string `json:"picurl"`
	Pictag string `json:"pictag"`
}

func NewPicInfo(pc Pictures) {
	sqlstr := `INSERT INTO pictureLibrary
	(noteID, picUrl, picTag)
	VALUES
	(?,?,?)`
	ret, err := db.Exec(sqlstr, pc.NoteId, pc.Picurl, pc.Pictag)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, affected rows:%d\n", theID)
}
