package mysqldb

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type NoteInfo struct {
	NoteId    int    `JSON:"NoteId"`
	NoteTitle string `JSON:"NoteTitle"`
	NoteCover string `JSON:"NoteCover"`
}

// 查询多条数据示例
func QueryNoteDemo() []NoteInfo {
	var Notes []NoteInfo
	sqlStr := "select noteId, title, cover from noteInfo where noteId > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt NoteInfo
		err := rows.Scan(&nt.NoteId, &nt.NoteTitle, &nt.NoteCover)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Notes = append(Notes, nt)
	}
	return Notes
}
