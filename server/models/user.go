package models

import "time"

type UserInfo struct {
	UserAccount  int       `JSON:"userAc"`
	UserName     string    `JSON:"username"`
	Password     string    `JSON:"password"`
	Gender       string    `JSON:"gender"`
	Portrait     string    `JSON:"portrait"`
	Introduction string    `JSON:"userName"`
	Birthday     time.Time `JSON:"birthday"`
	RegistTime   time.Time `JSON:"registTime"`
	FansNum      int       `JSON:"fansNum"`      // 粉丝数
	NoteNum      int       `JSON:"noteNum"`      // 笔记数
	CollectNum   int       `JSON:"collectNum"`   //收藏数
	FollowNum    int       `JSON:"followNum"`    //关注数
	CollectedNum int       `JSON:"collectedNum"` // 被收藏数量
	LikedNum     int       `JSON:"likedNum"`     // 被点赞数量
	PhoneNumber  string    `JSON:"phoneNumber"`
	Mail         string    `JSON:"mail"`
}
