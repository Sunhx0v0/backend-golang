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

func UserInfoDB(id int) UserInfo { // 从数据库获得用户信息
	sqlStr := `select userAccount, userName, gender, portrait, introduction, fansNum, noteNum, collectNum, followNum, collectedNum, likedNum, phoneNumber, mail, password,birthday,registTime
	from userinfo 
	where userAccount = ?`
	var ui UserInfo
	var bd, rt string
	err := db.QueryRow(sqlStr, id).Scan(&ui.UserID, &ui.UserName, &ui.Gender, &ui.Portrait, &ui.Introduction, &ui.FansNum, &ui.NoteNum, &ui.CollectNum, &ui.FollowNum, &ui.CollectedNum, &ui.LikedNum, &ui.PhoneNumber, &ui.Mail, &ui.Password, &bd, &rt)
	ui.Birthday = bd
	ui.RegistTime = rt
	// 如果要使用time.Time()类型的birthday/registTime，则将上面两句改成下面两句
	// ui.Birthday, _ = time.Parse("2006-01-02 15:04:05", bd)
	// ui.RegistTime, _ = time.Parse("2006-01-02 15:04:05", rt)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ui
	}
	fmt.Println(ui.UserID, ui.UserName, ui.Gender, ui.Portrait, ui.Introduction, ui.Birthday, ui.RegistTime, ui.FansNum, ui.NoteNum, ui.CollectNum, ui.FollowNum, ui.CollectedNum, ui.LikedNum, ui.PhoneNumber, ui.Mail, ui.Password)
	return ui
}

// 从数据库获得某用户发布的笔记信息
func NoteInfoDB(id int) []Notes {
	var notes []Notes
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.userName, u.portrait
	from noteInfo n, userInfo u
	where creatorAccount = ? and n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return notes
	}
	defer rows.Close()
	for rows.Next() {
		var nt Notes
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.CreatorName, &nt.Portrait)
		fmt.Println(nt.NoteID, nt.Title, nt.Cover, nt.CreatorID, nt.LikedNum, nt.CreatorName, nt.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return notes
		}
		notes = append(notes, nt)
	}
	return notes
}

// 从数据库获得某用户收藏的笔记信息
func CollectInfoDB(id int) []Notes {

	var collects []Notes
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.userName, u.portrait
	from noteInfo n, userInfo u, collectTable c
	where c.userAct = ? and c.collectNoteId=n.noteId and n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return collects
	}
	defer rows.Close()
	for rows.Next() {
		var ct Notes
		err := rows.Scan(&ct.NoteID, &ct.Title, &ct.Cover, &ct.CreatorID, &ct.LikedNum, &ct.CreatorName, &ct.Portrait)
		fmt.Println(ct.NoteID, ct.Title, ct.Cover, ct.CreatorID, ct.LikedNum, ct.CreatorName, ct.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return collects
		}
		collects = append(collects, ct)
	}
	return collects
}

func LikeInfoDB(id int) []Notes { // 从数据库获得用户信息

	var collects []Notes
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.userName, u.portrait
	from noteInfo n, userInfo u, favorTable c
	where c.userAct = ? and c.favorNoteId=n.noteId and n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return collects
	}
	defer rows.Close()
	for rows.Next() {
		var ct Notes
		err := rows.Scan(&ct.NoteID, &ct.Title, &ct.Cover, &ct.CreatorID, &ct.LikedNum, &ct.CreatorName, &ct.Portrait)
		fmt.Println(ct.NoteID, ct.Title, ct.Cover, ct.CreatorID, ct.LikedNum, ct.CreatorName, ct.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return collects
		}
		collects = append(collects, ct)
	}
	return collects
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
