package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type LoginInfo struct {
	Password string `json:"password"` // 密码
	UserID   int64  `json:"userId"`   // 用户编号
	UserName string `json:"userName"` // 用户名
}

type UserClaim struct {
	UserName string
	Claims   []LoginInfo
}

type userInfo struct {
	userId       int       `JSON:"UserId"`
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

func CheckUser(userName, password string) bool {
	//用户的登录信息
	var buser LoginInfo
	sqlstr := "select userAccount from userInfo where userName=? and password=?"
	err := db.QueryRow(sqlstr, userName, password).Scan(&buser.UserID)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	if buser.UserID > 0 {
		// fmt.Print(buser.UserID)
		return true
	}

	return false
}
