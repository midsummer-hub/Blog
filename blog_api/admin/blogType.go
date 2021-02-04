package admin

import (
	"blog/blog_api/service"
	"blog/blog_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//查询类型列表
func TypeList(c *gin.Context) {
	var page utils.Page
	err := c.ShouldBind(&page)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
		res.Json(c)
		return
	}
	blogType := &service.BlogType{}
	page.Total = blogType.Count()
	result := blogType.FindList(&page)
	res := &utils.Response{Code: 0, Msg: "", Data: result, Count: page.Total}
	res.Json(c)
}

//查询一个类型
func TypeOne(c *gin.Context)  {
	var blogType service.BlogType
	err := c.BindJSON(&blogType)
	if err != nil {
		res := &utils.Response{Code: 1000,Msg: "数据格式错误"}
		res.Json(c)
		return
	}
	result := blogType.FindOne()
	res := &utils.Response{Code: 0,Msg: "",Data: result}
	res.Json(c)
}

//查询所有类型
func TypeAll(c *gin.Context)  {
	blogType := &service.BlogType{}
	result := blogType.FindAll()
	res := &utils.Response{Code: 0,Msg: "",Data: result}
	res.Json(c)
}

//保存&更新博客类型
func BlogTypeSave(c *gin.Context)  {
	var blogTyle service.BlogType
	err := c.ShouldBind(&blogTyle)
	if err != nil{
		res := &utils.Response{Code: 1000,Msg: "数据格式错误"}
		res.Json(c)
		return
	}

	var result *gorm.DB
	if blogTyle.Id <= 0 {
		result = blogTyle.Insert()
	}else{
		result = blogTyle.Update()
	}
	if result.Error != nil {
		res := &utils.Response{Code: 1000,Msg: "提交错误"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0,Msg: ""}
	res.Json(c)
}

//删除博客类型
func BlogTypeDel(c *gin.Context)  {
	var blogTyle service.BlogType
	err := c.ShouldBind(&blogTyle)
	if err != nil{
		res := &utils.Response{Code: 1000,Msg: "数据格式错误"}
		res.Json(c)
		return
	}

	blog := new(service.Blog)
	blog.TypeId = blogTyle.Id
	count := blog.FindByTypeCount()
	if count > 0 {
		res := &utils.Response{Code: 1001,Msg: "此类型下有博文，请先删除相关博文"}
		res.Json(c)
		return
	}
	result := blogTyle.Delete()
	if result.Error != nil {
		res := &utils.Response{Code: 1000,Msg: "提交错误"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0,Msg: ""}
	res.Json(c)
}