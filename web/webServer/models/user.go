package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type userInfo struct {
	userAccount  int       `JSON:"UserAc"`
	userName     string    `JSON:"UserName"`
	password     string    `JSON:"Password"`
	gender       string    `JSON:"Gender"`
	portrait     string    `JSON:"Portrait"`
	introduction string    `JSON:"UserName"`
	birthday     time.Time `JSON:"Birthday"`
	registTime   time.Time `JSON:"RegistTime"`
	fansNum      int       `JSON:"FansNum"`      // 粉丝数
	noteNum      int       `JSON:"NoteNum"`      // 笔记数
	collectNum   int       `JSON:"CollectNum"`   //收藏数
	followNum    int       `JSON:"FollowNum"`    //关注数
	collectedNum int       `JSON:"CollectedNum"` // 被收藏数量
	likedNum     int       `JSON:"LikedNum"`     // 被点赞数量
	phoneNumber  string    `JSON:"PhoneNumber"`
	mail         string    `JSON:"Mail"`
}

func GetUserInfo(a int) {

}
