package admin

import (
	"blog/blog_api/cache"
	"blog/blog_api/middleware"
	"blog/blog_api/service"
	"blog/blog_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

//登录
func Login(c *gin.Context)  {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		res := &utils.Response{Code: 1000,Msg: "数据格式错误"}
		res.Json(c)
		return
	}
	//登录
	result := blogger.Login()
	if result == nil {
		res := &utils.Response{Code: 1001,Msg: "用户不存在"}
		res.Json(c)
		return
	}

	if result.Password != utils.Md5(blogger.Password) {
		res := &utils.Response{Code: 1001,Msg: "用户名密码错误"}
		res.Json(c)
		return
	}
	key := strconv.Itoa(time.Now().Nanosecond())

	token := &middleware.TokenData{
		Username:	utils.Md5(blogger.Password),
		Password:	utils.Md5(blogger.Username),
		Key:		key,
	}
	tokenKey := token.SetToken()
	//将token存入缓存
	utils.Cache.Set(tokenKey,token,cache.DefaultExpiration)
	res := &utils.Response{Code: 0,Msg: "",Data: tokenKey}
	res.Json(c)
}
//退出

func Logout(c *gin.Context)  {
	token := c.GetHeader("token")
	utils.Cache.Delete(token)
	res := &utils.Response{Code: 0,Msg: ""}
	res.Json(c)
}


//查找博主信息
func FindBlogger(c *gin.Context)  {
	var blogger service.Blogger
	result := blogger.Find()
	res := &utils.Response{Code: 0,Msg: "",Data: result}
	res.Json(c)
}

//修改个人博主信息
func BloggerUpdateInfo(c *gin.Context)  {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		res := &utils.Response{Code: 1000,Msg: "数据格式错误"}
		res.Json(c)
		return
	}
	var result *gorm.DB
	if blogger.Id <= 0 {
		result = blogger.Insert()
	}else{
		result = blogger.UpdateInfo()
	}
	if result.Error != nil {
		res := &utils.Response{Code: 1000,Msg: "提交错误"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0,Msg: ""}
	res.Json(c)
}

//修改博主密码
func BloggerUpdatePassword(c *gin.Context)  {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		res := &utils.Response{Code: 1000,Msg: "数据格式错误"}
		res.Json(c)
		return
	}
	if blogger.Password != "" {
		blogger.Password = utils.Md5(blogger.Password)
	}
	var result *gorm.DB
	if blogger.Id <= 0 {
		result = blogger.Insert()
	}else{
		result = blogger.UpdatePassword()
	}
	if result.Error != nil {
		res := &utils.Response{Code: 1000,Msg: "提交错误"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0,Msg: ""}
	res.Json(c)
}