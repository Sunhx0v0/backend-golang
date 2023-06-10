package models

import "time"

type NoteInfo struct {
	Noteid         int       `JSON:"noteid"`
	CreatorAccount int       `JSON:"creatorAccount"`
	Title          string    `JSON:"title"`
	Body           string    `JSON:"body"`       // 正文
	NumOfPic       int       `JSON:"numOfPic"`   // 图片数量
	Cover          string    `JSON:"cover"`      //封面图片
	CreateTime     time.Time `JSON:"createTime"` //创建时间
	UpdateTime     time.Time `JSON:"updateTime"` //最近更新
	Tag            string    `JSON:"tag"`        // 粉丝数
	Location       string    `JSON:"location"`   // 笔记数
	AtUserid       string    `JSON:"atUserid"`   //收藏数
}
