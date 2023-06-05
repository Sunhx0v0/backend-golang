package admin

import "github.com/gin-gonic/gin"

func SetupAdminRoutes(r *gin.Engine) {
	hostRegister(r)
}
