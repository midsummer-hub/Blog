package admin

import (
	"blog/blog_api/service"
	"blog/blog_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//返回博客列表
func BlogList(c *gin.Context) {
	var page utils.Page
	err := c.ShouldBind(&page)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
		res.Json(c)
		return
	}
	blog := new(service.Blog)
	page.Total = blog.Count()

	result, err := blog.FindList(&page)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: err.Error()}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: "", Data: result, Count: page.Total}
	res.Json(c)

}

//查找一个博客
func BlogOne(c *gin.Context) {
	var blog service.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
		res.Json(c)
		return
	}

	result := blog.FindOne()
	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

//保存博客内容
func BlogSave(c *gin.Context) {
	var blog service.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
		res.Json(c)
		return
	}

	var result *gorm.DB

	if blog.Id <= 0 {
		blog.AddTime = utils.GetDate(utils.DateFormat)
		result = blog.Insert()
	} else {
		blog.UpdateTime = utils.GetDate(utils.DateFormat)
		result = blog.Update()
	}

	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "提交出错"}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)

}

//删除博客
func BlogDelete(c *gin.Context) {
	var blog service.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
		res.Json(c)
		return
	}

	result := blog.Delete()

	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "提交出错"}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}
