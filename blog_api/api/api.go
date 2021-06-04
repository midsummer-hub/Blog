package api

import (
	"blog/blog_api/service"
	"blog/blog_api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//查找博主信息
func FindBlogger(c *gin.Context)  {
	var blogger service.Blogger

	result := blogger.Find()
	result.Password = ""
	res := &utils.Response{Code: 0,Msg: "",Data: result}
	res.Json(c)
}

//查找博客分类数量
func FindType(c *gin.Context)  {
	blogType := &service.BlogType{}
	result := blogType.FindTypeCount()
	res := &utils.Response{Code: 0,Msg: "",Data: result}
	res.Json(c)
}

//查找博客列表
func BlogList(c *gin.Context)  {
	json := make(map[string]interface{})
	err := c.ShouldBind(&json)
	if err != nil {
		res := &utils.Response{Code: 1000,Msg: "数据格式错误"}
		res.Json(c)
		return
	}
	blog := new(service.Blog)
	//字符串转int
	page,_:= strconv.Atoi(utils.StrVal(json["page"]))
	size,_:= strconv.Atoi(utils.StrVal(json["size"]))
	pageVo := &utils.Page{Page: page,Size: size,Total: blog.Count()}
	typeId,err := strconv.Atoi(utils.StrVal(json["type_id"]))

	if err == nil {
		blog.TypeId = typeId
	}
	//查询博客列表
	result, err := blog.FindList(pageVo)
	if err != nil {
		res := &utils.Response{Code: 1000,Msg: err.Error()}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0,Msg: "",Data: result,Count:pageVo.Total}
	res.Json(c)
}

//查询博客内容
func FindBlog(c *gin.Context)  {
	var blog service.Blog
	//绑定博客 ID
	err := c.BindJSON(&blog)
	if err != nil {
		res := &utils.Response{Code: 1000,Msg: "数据格式错误"}
		res.Json(c)
		return
	}
	//更新点击次数
	blog.UpdateClick()
	//根据博客ID查询类型名
	result := blog.FindtOneTypeName()
	//上一条
	last := blog.FindLastOne()
	//下一条
	next := blog.FindNextOne()
	//查找评论
	comments := blog.FindCommentByBlog()
	Map := make(map[string]interface{})
	Map["last"] = last
	Map["next"] = next
	Map["blog"] = result
	Map["comments"] = comments
	res := &utils.Response{Code: 0,Msg: "",Data: Map}
	res.Json(c)
}

//提交评论
func Comment(c *gin.Context)  {
	var comment service.Comment
	//ID内容
	err := c.BindJSON(&comment)
	if err != nil {
		res := &utils.Response{Code: 1000,Msg: "数据格式test错误"}
		res.Json(c)
		return
	}
	//IP
	comment.Ip = c.ClientIP()
	//添加时间
	comment.AddTime = utils.GetDate(utils.DateFormat)
	comment.Insert()

	blog := &service.Blog{Id:comment.BlogId}
	blog.UpdateReplay()
	res := &utils.Response{Code: 0,Msg: ""}
	res.Json(c)
}


