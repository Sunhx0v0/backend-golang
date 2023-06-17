package web

import (
	"github.com/gin-gonic/gin"
)

func SetupWebRoutes(r *gin.Engine) {
	userRegister(r)
	web_testRegister(r)
	exploreRegister(r)
}