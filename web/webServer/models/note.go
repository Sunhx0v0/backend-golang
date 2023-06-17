package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// type noteInfo struct {
// 	NoteId    int    `JSON:"NoteId"`
// 	NoteTitle string `JSON:"NoteTitle"`
// 	NoteCover string `JSON:"NoteCover"`
// }

// var Notes []noteInfo

type Note struct {
	Cover       string `json:"cover"`
	CreatorID   int    `json:"creatorId"`
	CreatorName string `json:"creatorName"` // 作者编号
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
func GetBriefNtInfo() (notes []Note) {
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.portrait, u.userName
	from noteInfo n,userInfo u
	where n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
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
			return
		}
		fmt.Print(nt.NoteID)
		notes = append(notes, nt)
	}
	return
}

// 获取特定内容的笔记
func GetSpBriefNtInfo(keyword string) (notes []Note) {
	sqlStr := `SELECT n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.portrait, u.userName
	FROM noteInfo n, userInfo u
	WHERE n.creatorAccount = u.userAccount AND (n.tag = ? OR n.title LIKE CONCAT('%',#{keyword},'%'))`
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

// 获取用户关注的人的所有笔记的简要信息
func GetFlwedNotes(userId int) (notes []Note, ok bool) {
	sqlStr := `SELECT n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.portrait, u.userName 
	FROM noteInfo n, userInfo u ,followTable f
	WHERE f.userAct=? AND f.followAct=n.creatorAccount AND n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr, userId)
	if err != nil {
		fmt.Printf("关注的人的笔记query failed, err:%v\n", err)
		return nil, false
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait, &nt.CreatorName)
		if err != nil {
			fmt.Printf("关注的人的笔记scan failed, err:%v\n", err)
			return nil, false
		}
		fmt.Println(nt.NoteID)
		notes = append(notes, nt)
	}
	return notes, true
}

// 存入新上传的笔记信息
func NewNoteInfo(nn DetailNote) (int, bool) {
	sqlstr := `INSERT INTO noteInfo
	(creatorAccount, cover, title, body, createTime, updateTime, tag, location, atUserId, likeNum)
	VALUES
	(?,?,?,?,?,?,?,?,?,?)`
	ret, err := db.Exec(sqlstr, nn.CreatorID, nn.Cover, nn.Title, nn.Body, nn.CreateTime, nn.UpdateTime, nn.Tag, nn.Location, nn.AtUserID, nn.LikedNum)
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

type detailNote struct {
	NoteInfo   DetailNote `JSON:"noteInfo"`
	PicsOfNote []string   `JSON:"pictures"`
}

// 查询特定笔记
func SpecificNote(noteid int) detailNote {
	var N detailNote
	//先找笔记信息
	sqlStr1 := "select Noteid,CreatorAccount,Title,Body,NumOfPic,Cover,CreateTime,UpdateTime,Tag,Location,AtUserid from noteInfo where noteId = ?"
	rows, err := db.Query(sqlStr1, noteid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		var err detailNote
		return err
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var createTimestring string
		var updateTimestring string
		err := rows.Scan(&N.NoteInfo.NoteID, &N.NoteInfo.CreatorID, &N.NoteInfo.Title, &N.NoteInfo.Body, &N.NoteInfo.Picnum, &N.NoteInfo.Cover, &createTimestring, &updateTimestring, &N.NoteInfo.Tag, &N.NoteInfo.Location, &N.NoteInfo.AtUserID)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			var err detailNote
			return err
		}
		N.NoteInfo.CreateTime, _ = time.Parse("2000-01-01 24:00:00", createTimestring)
		N.NoteInfo.UpdateTime, _ = time.Parse("2000-01-01 24:00:00", updateTimestring)
	}
	//再找图片信息
	sqlStr2 := "select picUrl from pictureLibrary where noteID = ?"
	rows2, err := db.Query(sqlStr2, noteid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		var err detailNote
		return err
	}
	for rows2.Next() {
		var picurl string
		err := rows.Scan(&picurl)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			var err detailNote
			return err
		}
		N.PicsOfNote = append(N.PicsOfNote, picurl)
	}

	// 关闭rows释放持有的数据库链接
	defer rows2.Close()

	return N
}

// 修改笔记被收藏数
func ChangeNoteCollects(noteId, option int) {
	var sqlstr string
	addnum := `UPDATE noteInfo set collectNum =collectNum+1 WHERE noteId = ?`
	reducenum := `UPDATE noteInfo set collectNum =collectNum-1 WHERE noteId = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, noteId)
	if err != nil {
		fmt.Printf("笔记收藏数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("笔记收藏数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("笔记收藏数修改编号：%d\n", n)
}
