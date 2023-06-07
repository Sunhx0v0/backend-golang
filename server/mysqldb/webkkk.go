package mysqldb

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 定义结构
type User struct {
	Username string `json:"username"` // 定制结构体字段为username，避开一定要大写才能共享的约束（代码大写，显示小写）
	Password string `json:"password"`
}

// 查询全部数据
func Findalldata(ctx *gin.Context) {
	err := InitDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	var users []User
	sqlStr := "select username, password from easy"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt User
		err := rows.Scan(&nt.Username, &nt.Password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		users = append(users, nt)
	}
	// var usrs = gin.H{
	// 	"username": users.username,
	// 	"password": users.password,
	// }
	ctx.JSON(http.StatusOK, users)
}

// // 查询全部数据
// func findAll(ctx *gin.Context) {
// 	url := "mongodb://127.0.0.1"
// 	session, err := mgo.Dial(url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	session.SetMode(mgo.Monotonic, true)
// 	c := session.DB("db_go").C("user")

// 	// 查找全部
// 	usrs := make([]user, 10)
// 	// 查找全部
// 	err = c.Find(nil).All(&usrs)
// 	ctx.JSON(http.StatusOK, usrs)
// }

// 查询多条数据示例
// func pueryNoteDemo() []NoteInfo {
// 	var Notes []NoteInfo
// 	sqlStr := "select noteId, title, cover from noteInfo where noteId > ?"
// 	rows, err := db.Query(sqlStr, 0)
// 	if err != nil {
// 		fmt.Printf("query failed, err:%v\n", err)
// 		return nil
// 	}
// 	// 关闭rows释放持有的数据库链接
// 	defer rows.Close()

// 	// 循环读取结果集中的数据
// 	for rows.Next() {
// 		var nt NoteInfo
// 		err := rows.Scan(&nt.NoteId, &nt.NoteTitle, &nt.NoteCover)
// 		if err != nil {
// 			fmt.Printf("scan failed, err:%v\n", err)
// 			return nil
// 		}
// 		Notes = append(Notes, nt)
// 	}
// 	return Notes
// }
