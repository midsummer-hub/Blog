package middleware

import (
	"blog/blog_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := utils.Log()
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()

		//执行时间
		latencyTime := endTime.Sub(startTime) / time.Millisecond

		//请求方式
		reMethod := c.Request.Method

		//请求路由
		reqUri := c.Request.RequestURI

		//请求头
		header := c.Request.Header

		//协议
		proto := c.Request.Proto

		//状态码
		statusCode := c.Writer.Status()

		//请求IP
		clientIP := c.ClientIP()

		err := c.Err()

		body,_ := ioutil.ReadAll(c.Request.Body)

		logger.WithFields(logrus.Fields{
			"status_code":	statusCode,
			"latency_Time":	latencyTime,
			"clientIP":		clientIP,
			"reMethod":		reMethod,
			"reqUri":		reqUri,
			"header":		header,
			"proto":		proto,
			"err":			err,
			"body":			body,
		}).Info()


	}
}