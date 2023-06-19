package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Regist struct { // 注册用的结构体
	Birthday     string  `json:"birthday"`
	Gender       string  `json:"gender "`      // 性别
	Introduction *string `json:"introduction"` // 简介
	Mail         string  `json:"mail"`
	Password     string  `json:"password"`
	PhoneNumber  string  `json:"phoneNumber"`
	Portrait     string  `json:"portrait"` // 头像
	UserName     string  `json:"userName "`
}

func IsTelephoneExists(PhoneNumber string) bool { // 查找手机号是否存在
	userID := 0
	sqlStr := `select userAccount
	from userinfo 
	where phoneNumber = ?`
	err := db.QueryRow(sqlStr, PhoneNumber).Scan(&userID)
	if err != nil {
		return false
	}
	return true
}

func CreateUser(registInfo Regist, userID int, registTime string) { // 创建用户
	sqlStr := `insert into userinfo(userName, password, gender,portrait,introduction,birthday,phoneNumber,mail,userAccount,registTime) values(?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, registInfo.UserName, registInfo.Password, registInfo.Gender, registInfo.Portrait, registInfo.Introduction, registInfo.Birthday, registInfo.PhoneNumber, registInfo.Mail, userID, registTime)
	if err != nil {
		return
	}
	return
}

func SecretCorrect(PhoneNumber string, password string) bool { // 判断密码是否正确
	sqlStr := `select password
	from userinfo 
	where phoneNumber = ?`
	var pwd string
	err := db.QueryRow(sqlStr, PhoneNumber).Scan(&pwd)
	if err != nil {
		return false
	}
	if pwd != password {
		return false
	}
	return true
}

func SelectAll(phone string) Regist {
	var rg Regist
	sqlStr := `select userName, password, gender,portrait,introduction,birthday,phoneNumber,mail
	from userinfo 
	where phoneNumber = ?`
	err := db.QueryRow(sqlStr, phone).Scan(&rg.UserName, &rg.Password, &rg.Gender, &rg.Portrait, &rg.Introduction, &rg.Birthday, &rg.PhoneNumber, &rg.Mail)
	if err != nil {
		return rg
	}
	return rg
}
