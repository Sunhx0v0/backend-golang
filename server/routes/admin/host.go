package admin

// import (
// 	"fmt"
// 	"net/http"
// 	"server/mysqldb"

// 	"github.com/gin-gonic/gin"
// )

// func hostRegister(r *gin.Engine) {
// 	r.GET("/host", hostHandler)
// }

// func hostHandler(c *gin.Context) {
// 	err := mysqldb.InitDB() // 调用输出化数据库的函数
// 	if err != nil {
// 		fmt.Printf("init db failed,err:%v\n", err)
// 		return
// 	}
// 	// 获取查询结果
// 	notes := mysqldb.QueryNoteDemo()
// 	// 将查询结果发送到前端
// 	c.HTML(http.StatusOK, "host.html", gin.H{
// 		"title": "标题",
// 		"nt":    notes,
// 	})
// }
