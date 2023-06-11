package models

import (
	"fmt"

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
	Cover       string `json:"cover"`
	LikedNum    int64  `json:"likedNum"` // 点赞数
	NoteID      int64  `json:"noteId"`   // 笔记编号
	Title       string `json:"title"`
	CreatorID   int64  `json:"creatorID"`
	CreatorName string `json:"creatorName"` // 作者姓名
	Portrait    string `json:"portrait"`
}

// 点赞的笔记
type Likes struct {
	Cover       string `json:"cover"`
	LikedNum    int64  `json:"likedNum"` // 点赞数
	NoteID      int64  `json:"noteId"`   // 笔记编号
	Title       string `json:"title"`
	CreatorID   int64  `json:"creatorID"`
	CreatorName string `json:"creatorName"` // 作者姓名
	Portrait    string `json:"portrait"`
}

// 发布的笔记
type Notes struct {
	Cover       string `json:"cover"`
	LikedNum    int64  `json:"likedNum"` // 点赞数
	NoteID      int64  `json:"noteId"`   // 笔记编号
	Title       string `json:"title"`
	CreatorID   int64  `json:"creatorID"`
	CreatorName string `json:"creatorName"` // 作者姓名
	Portrait    string `json:"portrait"`
}

// 用户基本信息
type UserInfo struct {
	Birthday     string `json:"birthday"`
	CollectedNum int64  `json:"collectedNum"` // 被收藏的笔记数
	CollectNum   int64  `json:"collectNum"`   // 收藏数
	FansNum      int64  `json:"fansNum"`
	FollowNum    int64  `json:"followNum"`    // 关注数
	Gender       string `json:"gender "`      // 性别
	Introduction string `json:"introduction"` // 简介
	LikedNum     int64  `json:"likedNum"`     // 被点赞数量
	Mail         string `json:"mail"`
	NoteNum      int64  `json:"noteNum "`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phoneNumber"`
	Portrait     string `json:"portrait"` // 头像
	RegistTime   string `json:"registTime"`
	UserID       int64  `json:"userAccount"`
	UserName     string `json:"userName "`
}

func UserInfoDB(id int) []UserInfo { // 从数据库获得用户信息
	// var ui UserInfo
	// sqlStr := `select userAccount, userName, gender, portrait, introduction, fansNum, noteNum, collectNum, followNum, collectedNum, likedNum, phoneNumber, mail, password
	// from userinfo
	// where userAccount = ?`
	// err := db.QueryRow(sqlStr, id).Scan(&ui.UserID, &ui.UserName, &ui.Gender, &ui.Portrait, &ui.Introduction, &ui.FansNum, &ui.NoteNum, &ui.CollectNum, &ui.FollowNum, &ui.CollectedNum, &ui.LikedNum, &ui.PhoneNumber, &ui.Mail, &ui.Password)
	// ui.Birthday = time.Now().Format("2006-01-02 15:04:05")
	// ui.RegistTime = time.Now().Format("2006-01-02 15:04:05")
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)
	// 	return uui
	// }
	// uui = append(uui, ui)
	// return uui

	var uui []UserInfo
	sqlStr := `select userAccount, userName, gender, portrait, introduction, fansNum, noteNum, collectNum, followNum, collectedNum, likedNum, phoneNumber, mail, password
	from userinfo
	where userAccount = ?`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		return uui
	}
	defer rows.Close()
	for rows.Next() {
		var ui UserInfo
		err := rows.Scan(&ui.UserID, &ui.UserName, &ui.Gender, &ui.Portrait, &ui.Introduction, &ui.FansNum, &ui.NoteNum, &ui.CollectNum, &ui.FollowNum, &ui.CollectedNum, &ui.LikedNum, &ui.PhoneNumber, &ui.Mail, &ui.Password)
		if err != nil {
			return uui
		}
		// ui.Birthday = time.Now().Format("2006-01-02 15:04:05")
		// ui.RegistTime = time.Now().Format("2006-01-02 15:04:05")
		uui[0] = ui
		//uui = append(uui, ui)
	}
	return uui

	// var ui UserInfo
	// ui.UserID = 10001
	// ui.UserName = "zahgsad"
	// ui.Gender = "sha"
	// ui.Portrait = "/"
	// ui.Introduction = "jntm"
	// ui.Birthday = time.Now().Format("2006-01-02 15:04:05")
	// ui.RegistTime = time.Now().Format("2006-01-02 15:04:05")
	// ui.FansNum = 666
	// ui.NoteNum = 777
	// ui.CollectNum = 888
	// ui.FollowNum = 999
	// ui.CollectedNum = 1010
	// ui.LikedNum = 115
	// ui.PhoneNumber = "1233663"
	// ui.Mail = "1235"
	// ui.Password = "123456"
	// uui = append(uui, ui)
	// return uui
}

func NoteInfoDB(id int) (notes []Notes) { // 从数据库获得用户信息
	// sqlStr := `select n.noteId, n.title, n.cover,n.creatorAccount,n.likeNum, u.portrait
	// from noteInfo n,userInfo u
	// where n.creatorAccount = u.userAccount`
	// rows, err := db.Query(sqlStr)
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)
	// 	return
	// }
	// // 关闭rows释放持有的数据库链接
	// defer rows.Close()

	// // 循环读取结果集中的数据
	// for rows.Next() {
	// 	var nt Notes
	// 	err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait)
	// 	if err != nil {
	// 		fmt.Printf("scan failed, err:%v\n", err)
	// 		return
	// 	}
	// 	fmt.Print(nt.NoteID)
	// 	notes = append(notes, nt)
	// }
	// return
	var nt Notes
	nt.NoteID = 10001
	nt.Title = "l就别"
	nt.Cover = "站发给"
	nt.CreatorID = 10001
	nt.LikedNum = 500
	nt.Portrait = "/"
	nt.CreatorName = "张菲"
	notes = append(notes, nt)
	return
}

func CollectInfoDB(id int) (notes []Collects) { // 从数据库获得用户信息
	// sqlStr := `select n.noteId, n.title, n.cover,n.creatorAccount,n.likeNum, u.portrait
	// from noteInfo n,userInfo u
	// where n.creatorAccount = '"+id+"' and u.userAccount = '"+id+"'`
	// rows, err := db.Query(sqlStr)
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var nt Collects
	// 	err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait)
	// 	if err != nil {
	// 		fmt.Printf("scan failed, err:%v\n", err)
	// 		return
	// 	}
	// 	notes = append(notes, nt)
	// }
	// return
	var nt Collects
	nt.NoteID = 10001
	nt.Title = "l就别"
	nt.Cover = "站发给"
	nt.CreatorID = 10001
	nt.LikedNum = 500
	nt.Portrait = "/"
	nt.CreatorName = "张菲"
	notes = append(notes, nt)
	return
}

func LikeInfoDB(id int) (notes []Likes) { // 从数据库获得用户信息
	// sqlStr := `select n.noteId, n.title, n.cover,n.creatorAccount,n.likeNum, u.portrait
	// from noteInfo n,userInfo u
	// where n.creatorAccount = '"+id+"' and u.userAccount = '"+id+"'`
	// rows, err := db.Query(sqlStr)
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var nt Likes
	// 	err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait)
	// 	if err != nil {
	// 		fmt.Printf("scan failed, err:%v\n", err)
	// 		return
	// 	}
	// 	notes = append(notes, nt)
	// }
	// return
	var nt Likes
	nt.NoteID = 10001
	nt.Title = "l就别"
	nt.Cover = "站发给"
	nt.CreatorID = 10001
	nt.LikedNum = 500
	nt.Portrait = "/"
	nt.CreatorName = "张菲"
	notes = append(notes, nt)
	return
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
