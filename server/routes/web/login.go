package web

import (
	"server/mysqldb"

	"github.com/gin-gonic/gin"
)

func loginRegister(r *gin.Engine) {
	// r.POST("/login", mysqldb.LoginHandler)
	r.POST("/mongo/login", mysqldb.LoginHandler)
}
