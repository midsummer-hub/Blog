package middleware

import (
	"blog/blog_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TokenData struct {
	Username string
	Password string
	Key 	 string
}

func (token *TokenData) SetToken() string{
	return utils.Md5(token.Username+token.Password+token.Key)
}

func Jwt() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			res := &utils.Response{
				Code:1100,
				Msg: "请求未携带token,无权访问",
			}
			res.Json(c)
			c.Abort()
			return
		}
		logrus.Debug("get Token:",token)
		data,fount := utils.Cache.Get(token)
		if fount == false {
			res := &utils.Response{
				Code:1100,
				Msg: "token找不到或已过期",
			}
			res.Json(c)
			c.Abort()
			return
		}
		tokenData := data.(*TokenData)

		b := token != tokenData.SetToken()
		if b {
			res := &utils.Response{
				Code:1100,
				Msg: "token认证错误",
			}
			res.Json(c)
			c.Abort()
			return
		}
		c.Set("token",data)
	}
}





































