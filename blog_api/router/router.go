package router

import (
	"blog/blog_api/admin"
	"blog/blog_api/api"
	"blog/blog_api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	//宕机时可以恢复
	router.Use(gin.Recovery())

	router.Static("/static","static")
	//跨域中间件
	router.Use(middleware.Cors())

	//日志中间件
	router.Use(middleware.Logger())


	register(router)

	return router
}

func register(router *gin.Engine)  {
	//api
	//查找博主
	router.POST("/blogger",api.FindBlogger)
	//查找博客类型
	router.POST("/blog/type",api.FindType)
	//查询博客列表
	router.POST("/blog/list",api.BlogList)
	//查询博客内容
	router.POST("/blog/show",api.FindBlog)
	//提交评论
	router.POST("/blog/comment",api.Comment)

	//admin
	//登录
	router.POST("/login",admin.Login)
	//退出
	router.POST("/logout",admin.Logout)
	jwt := router.Group("/admin",middleware.Jwt())
	{
		//博主信息
		jwt.POST("/blogger/find",admin.FindBlogger)
		//更新博主信息
		jwt.POST("/blogger/updateInfo",admin.BloggerUpdateInfo)
		//更新博主密码
		jwt.POST("/blogger/updatePassword",admin.BloggerUpdatePassword)

	}

}


























