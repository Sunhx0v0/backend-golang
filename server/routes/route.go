package routes

import (
	"server/routes/admin"
	"server/routes/web"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	admin.SetupAdminRoutes(r)
	web.SetupWebRoutes(r)
	// 一次加载所有html（不一定用得上）
	r.LoadHTMLGlob("./templates/**/*")
	return r
}
