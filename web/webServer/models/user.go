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

type UserClaim struct { // 登录验证
	UserName string
	Claims   []LoginInfo
}

// 收藏的笔记
type Collects struct {
	Cover     string `json:"cover"`
	LikedNum  int64  `json:"likedNum"` // 点赞数
	NoteID    int64  `json:"noteId"`   // 笔记编号
	Title     string `json:"title"`
	CreatorID int64  `json:"creatorID"`
	Portrait  int64  `json:"portrait"`
}

// 点赞的笔记
type Likes struct {
	Cover     string `json:"cover"`
	LikedNum  int64  `json:"likedNum"` // 点赞数
	NoteID    int64  `json:"noteId"`   // 笔记编号
	Title     string `json:"title"`
	CreatorID int64  `json:"creatorID"`
	Portrait  int64  `json:"portrait"`
}

// 发布的笔记
type Notes struct {
	Cover     string `json:"cover"`
	LikedNum  int64  `json:"likedNum"` // 点赞数
	NoteID    int64  `json:"noteId"`   // 笔记编号
	Title     string `json:"title"`
	CreatorID int64  `json:"creatorID"`
	Portrait  int64  `json:"portrait"`
}

// 用户基本信息
type UserInfo struct {
	userID       int       `JSON:"UserID"`
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

func UserInfoDB(id int) UserInfo { // 从数据库获得用户信息
	sqlStr := `select *
	from userinfo
	where userAccount = '"+id+"'`
	rows, err := db.Query(sqlStr)
	var ui UserInfo
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return ui
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		err := rows.Scan(&ui.userID, &ui.userName, &ui.gender, &ui.portrait, &ui.introduction, &ui.birthday, &ui.registTime, &ui.fansNum, &ui.noteNum, &ui.collectNum, &ui.followNum, &ui.collectedNum, &ui.likedNum, &ui.phoneNumber, &ui.mail, &ui.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return ui
		}
	}
	return ui
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
