package mysqldb

import "fmt"

func UpdateOneAttribute_int(id int, attribute string, variant int) {
	sqlstr := "UPDATE userInfo set ? = ? WHERE userAccount = ?"
	db.Exec(sqlstr, attribute, variant, id)
}

//临时
func ChangeUserFollows(userId, option int) {
	var sqlstr string
	addnum := `UPDATE userInfo set followNum =followNum+1 WHERE userAccount = ?`
	reducenum := `UPDATE userInfo set followNum =followNum-1 WHERE userACCount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("关注数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("关注数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("关注数修改编号：%d\n", n)
}

//临时
func ChangeUserFans(userId, option int) {
	var sqlstr string
	addnum := `UPDATE userInfo set fansNum =fansNum+1 WHERE userAccount = ?`
	reducenum := `UPDATE userInfo set fansNum =fansNum-1 WHERE userAccount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("粉丝数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("粉丝数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("粉丝数修改编号：%d\n", n)
}
