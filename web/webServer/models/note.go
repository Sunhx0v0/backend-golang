package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 要返回给前端的数据对应的结构体
type Note struct {
	Cover       string `json:"cover"`
	CreatorID   int    `json:"creatorId"`   // 作者编号
	CreatorName string `json:"creatorName"` // 作者姓名
	LikedNum    int    `json:"likedNum"`    // 点赞数
	NoteID      int    `json:"noteId"`      // 笔记编号
	Portrait    string `json:"portrait"`    // 头像
	Title       string `json:"title"`
}

// 笔记的详细信息
type DetailNote struct {
	NoteID     int       `json:"noteid"`    // 笔记编号
	CreatorID  int       `json:"creatorId"` // 作者编号
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Picnum     int       `json:"picnum"`
	Cover      string    `json:"cover"`
	CreateTime time.Time `json:"createtime"`
	UpdateTime time.Time `json:"updatetime"`
	Tag        string    `json:"tag"`
	Location   string    `json:"location"`
	AtUserID   int       `json:"atuserid"`
	LikedNum   int       `json:"likedNum"` // 点赞数
}

// 获取笔记的封面标题等简要信息
func GetBriefNtInfo() (notes []Note, ok bool) {
	ok = true
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.portrait, u.userName
	from noteInfo n, userInfo u
	where n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		ok = false
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait, &nt.CreatorName)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			ok = false
			return
		}
		// fmt.Printf("笔记编号：%d ", nt.NoteID)
		notes = append(notes, nt)
	}
	return
}

// 获取特定内容的笔记
func GetSpBriefNtInfo(keyword string) (notes []Note) {
	sqlStr := `SELECT n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.portrait, u.userName 
	FROM noteInfo n, userInfo u 
	WHERE n.creatorAccount = u.userAccount AND ( n.tag=? OR n.title LIKE CONCAT('%',?,'%'))`
	rows, err := db.Query(sqlStr, keyword, keyword)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait, &nt.CreatorName)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Println(nt.NoteID)
		notes = append(notes, nt)
	}
	return
}

// 存入新上传的笔记信息
func NewNoteInfo(nn DetailNote) (int, bool) {
	sqlstr := `INSERT INTO noteInfo (creatorAccount, title, body, createTime, updateTime) VALUES (?,?,?,?,?)`
	ret, err := db.Exec(sqlstr, nn.CreatorID, nn.Title, nn.Body, nn.CreateTime, nn.UpdateTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return -1, false
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return -1, false
	}
	return int(theID), true
}

// 更新某笔记的信息
func ModifyNote(mn DetailNote) bool {
	sqlstr := `UPDATE noteInfo SET
	cover=?, title=?, body=?, createTime=?, updateTime=?, tag=?, location=?, atUserId=?
	WHERE
	noteId=?`
	ret, err := db.Exec(sqlstr, mn.Cover, mn.Title, mn.Body, mn.CreateTime, mn.UpdateTime, mn.Tag, mn.Location, mn.AtUserID, mn.NoteID)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	theID, err := ret.LastInsertId() // 修改的行数
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return false
	}
	fmt.Printf("笔记更新成功！，影响行数：%d\n", theID)
	return true
}

// 删除笔记
func DeleteNote(ntid int) bool {
	sqlstr := "DELETE FROM noteInfo WHERE noteID = ?"
	ret, err := db.Exec(sqlstr, ntid)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return false
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
	return true
}
