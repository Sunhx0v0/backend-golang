package mysqldb

import (
	"fmt"
	"server/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type NoteInfo models.NoteInfo
type PictureInfo models.PictureInfo

// ——————————————————————————————————————————————————————
// 查询特定笔记
type Note struct {
	NoteInfo   NoteInfo `JSON:"noteInfo"`
	PicsOfNote []string `JSON:"pictures"`
}

func SpecificNote(noteid int) Note {
	var N Note
	//先找笔记信息
	sqlStr1 := "select Noteid,CreatorAccount,Title,Body,NumOfPic,Cover,CreateTime,UpdateTime,Tag,Location,AtUserid from noteInfo where noteId = ?"
	rows, err := db.Query(sqlStr1, noteid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		var err Note
		return err
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var createTimestring string
		var updateTimestring string
		err := rows.Scan(&N.NoteInfo.Noteid, &N.NoteInfo.CreatorAccount, &N.NoteInfo.Title, &N.NoteInfo.Body, &N.NoteInfo.NumOfPic, &N.NoteInfo.Cover, &createTimestring, &updateTimestring, &N.NoteInfo.Tag, &N.NoteInfo.Location, &N.NoteInfo.AtUserid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			var err Note
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
		var err Note
		return err
	}
	for rows2.Next() {
		var picurl string
		err := rows.Scan(&picurl)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			var err Note
			return err
		}
		N.PicsOfNote = append(N.PicsOfNote, picurl)
	}

	// 关闭rows释放持有的数据库链接
	defer rows2.Close()

	return N
}

// ——————————————————————————————————————————————————————————————————
