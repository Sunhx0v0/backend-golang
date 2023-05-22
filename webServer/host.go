// 主页面
package main

import (
	"fmt"

	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

// 全局对象db
var db *sql.DB

type noteInfo struct {
	noteId    int
	noteTitle string
	noteCover string
}

var notes []noteInfo

// 初始化数据库
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/webData"
	// 不会校验账号密码是否正确
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt noteInfo
		err := rows.Scan(&nt.noteId, &nt.noteTitle, &nt.noteCover)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		notes = append(notes, nt)
		//fmt.Printf("id:%d name:%s age:%s\n", nt.noteId, nt.noteTitle, nt.noteCover)
	}
}

func hostFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "host.html", gin.H{
		"title": "标题",
		"nt":    notes,
	})
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	queryMultiRowDemo()

	router := gin.Default()

	//加载HTML文件
	router.LoadHTMLFiles("templates/host.html")
	router.GET("/host", hostFunc)
	router.Run(":8080")
}
