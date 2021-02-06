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

	//文件上传
	router.POST("/upload",admin.Upload)
	//文件上传编辑
	router.POST("/uploadEdit",admin.UploadEdit)
	jwt := router.Group("/admin",middleware.Jwt())
	{
		//博主信息
		jwt.POST("/blogger/find",admin.FindBlogger)
		//更新博主信息
		jwt.POST("/blogger/updateInfo",admin.BloggerUpdateInfo)
		//更新博主密码
		jwt.POST("/blogger/updatePassword",admin.BloggerUpdatePassword)

		//博客类型list
		jwt.POST("/blog/type/list",admin.TypeList)
		//保存&修改博客类型
		jwt.POST("/blog/type/save",admin.BlogTypeSave)
		//查询一个类型
		jwt.POST("/blog/type/one",admin.TypeOne)
		//查询所有类型
		jwt.POST("/blog/type/All",admin.TypeAll)
		//删除博客类型
		jwt.POST("/blog/type/delete",admin.BlogTypeDel)

		//博客文章list
		jwt.POST("/blog/list",admin.BlogList)
		//博客文章内容保存
		jwt.POST("/blog/save",admin.BlogSave)
		//博客文章删除
		jwt.POST("/blog/delete",admin.BlogDelete)
		//查找一篇文章博客
		jwt.POST("/blog/one",admin.BlogOne)

		//博客评论list
		jwt.POST("/comment/list",admin.CommentList)
		//博客评论状态审核
		jwt.POST("/comment/review",admin.CommentReview)
		//博客评论删除
		jwt.POST("/comment/delete",admin.CommentDelete)


	}

}


























