package routers

import (
	"webServer/middleware/cors"
	"webServer/middleware/webjwt"
	v1 "webServer/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	r.POST("/register", v1.Register) // 注册

	//使用Use（）方法向路由器添加一些中间件功能。
	r.Use(gin.Logger())   //第一个中间件函数是gin.Logger（），它记录对控制台或文件的HTTP请求和响应。它帮助开发人员调试和监控应用程序的行为。
	r.Use(gin.Recovery()) //第二个中间件函数是gin.Recovery（），它可以从请求处理过程中发生的任何死机中恢复。它确保服务器不会因意外错误而崩溃，并返回错误响应。
	// 使用CorsMiddleware()中间件来进行跨域连接
	r.Use(cors.CorsMiddleware()) //最后，它添加了一个名为CORS.CorsMiddleware（）的第三方CORS中间件。该中间件允许跨源资源共享（CORS），使运行在不同域上的客户端JavaScript应用程序能够访问

	// gin.SetMode(setting.RunMode)
	var userMiddleware = webjwt.GinJWTMiddlewareInit(&webjwt.Visitor{}) // 自定义的授权规则
	r.POST("/login", userMiddleware.LoginHandler)
	//404 handler
	r.NoRoute(userMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "Page not found",
		})
	})

	user := r.Group("/user")
	{
		// 刷新token
		user.GET("/refresh_token", userMiddleware.RefreshHandler)
	}

	// api := r.Group("/user")
	// api.Use(authMiddleware.MiddlewareFunc())
	// {
	// 	api.GET("/info", v1.GetUserInfo)
	// 	api.POST("/logout", v1.Logout)
	// }

	apiv1 := r.Group("/api/v1")
	//使用userAuthorizator中间件，只有user权限的用户才能获取到接口
	apiv1.Use(userMiddleware.MiddlewareFunc())
	{
		//获取笔记（全部）
		r.GET("/explore", v1.GetAllNotes)
		//获取特定笔记（搜索/标签）
		r.GET("/search/:keyword", v1.GetSpecificNotes)
		//获取关注人的笔记
		r.GET("/:userId/follow/notes", v1.GetFollowedNotes)
		//获取关注的人
		r.GET("/userId/follow")

		// 获取用户界面的信息
		r.GET("/:userId/PersonalView", v1.GetUserInfo)
		// 用户修改个人信息
		r.PUT("/:userId/PersonalView", v1.ModifyUserInfo)
		//上传笔记
		r.POST("/:userId/publish", v1.UploadNote)
		//用户删除笔记
		r.DELETE("/:userId/publish/:noteId", v1.DeleteNote)

		//加载某篇笔记的评论
		r.GET("/comment/:noteId", v1.GetComments)
		//发表评论
		r.POST("/comment/:noteId", v1.PostComment)
		//删除评论
		r.DELETE("/comment/:noteId", v1.CancleComment)

		//点赞某篇笔记
		r.POST("/explore/:noteId/like", v1.LikeNote)
		//取消点赞
		r.DELETE("/explore/:noteId/like", v1.CancelLike)
		//获取笔记详细内容
		r.GET("/explore/:noteid", v1.NoteDetailHandler)
		//收藏某篇笔记
		r.POST("/explore/:noteId/collect", v1.CollectNote)
		//取消收藏某篇笔记
		r.DELETE("/explore/:noteId/collect", v1.CancleCollect)

		// 关注用户
		r.POST("/:userId/PersonalView/follow", v1.FollowHandler)

		//获取评论消息列表
		r.GET("/messages/:userId/comments", v1.MsgGetComments)
		//把评论设置为已读
		r.PUT("/messages/:userId/comments/:commentId", v1.ChangeCommentState)
		//获取点赞消息列表
		r.GET("/messages/:userId/likes")
		//把点赞消息设置为已读
		r.PUT("/messages/:userId/likes/:fvId")
	}

	return r
}
