package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int 			`json:"code"`
	Msg string 			`json:"msg"`
	Data interface{} 	`json:"data,omitempty"`
	Count interface{} 	`json:"count,omitmepty"`
}

//返回json数据
func (res *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK,res)
	return
}