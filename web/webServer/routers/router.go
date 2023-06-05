package routers

import (
	"github.com/gin-gonic/gin"
	"webServer/middleware/cors"
	//"webServer/middleware/webjwt"
	"webServer/routers/api/v1"
)

func InitRouter() *gin.Engine {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(cors.CorsHandler())

	// 使用CorsMiddleware()中间件来进行跨域连接
	r.Use(cors.CorsMiddleware())

	//加载HTML文件
	r.LoadHTMLGlob("templates/host.html")
	r.GET("/explore", v1.HostFunc)

	// gin.SetMode(setting.RunMode)
	// var authMiddleware = jwt.GinJWTMiddlewareInit(&myjwt.AllUserAuthorizator{})
	// r.POST("/login", authMiddleware.LoginHandler)
	// //404 handler
	// r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
	// 	code := e.PAGE_NOT_FOUND
	// 	c.JSON(404, gin.H{"code": code, "message": e.GetMsg(code)})
	// })

	// auth := r.Group("/auth")
	// {
	// 	// Refresh time can be longer than token timeout
	// 	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	// }

	// api := r.Group("/user")
	// api.Use(authMiddleware.MiddlewareFunc())
	// {
	// 	api.GET("/info", v1.GetUserInfo)
	// 	api.POST("/logout", v1.Logout)
	// }

	// var adminMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.AdminAuthorizator{})
	// apiv1 := r.Group("/api/v1")
	// //使用AdminAuthorizator中间件，只有admin权限的用户才能获取到接口
	// apiv1.Use(adminMiddleware.MiddlewareFunc())
	// {
	// 	//vue获取table信息
	// 	apiv1.GET("/table/list", v2.GetArticles)
	// 	//获取标签列表
	// 	apiv1.GET("/tags", v1.GetTags)
	// 	//新建标签
	// 	apiv1.POST("/tags", v1.AddTag)
	// 	//更新指定标签
	// 	apiv1.PUT("/tags/:id", v1.EditTag)
	// 	//删除指定标签
	// 	apiv1.DELETE("/tags/:id", v1.DeleteTag)

	// 	//获取文章列表
	// 	apiv1.GET("/articles", v1.GetArticles)
	// 	//获取指定文章
	// 	apiv1.GET("/articles/:id", v1.GetArticle)
	// 	//新建文章
	// 	apiv1.POST("/articles", v1.AddArticle)
	// 	//更新指定文章
	// 	apiv1.PUT("/articles/:id", v1.EditArticle)
	// 	//删除指定文章
	// 	apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	// }

	// var testMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.TestAuthorizator{})
	// apiv2 := r.Group("/api/v2")
	// apiv2.Use(testMiddleware.MiddlewareFunc())
	// {
	// 	//获取文章列表
	// 	apiv2.GET("/articles", v2.GetArticles)
	// }

	return r
}
