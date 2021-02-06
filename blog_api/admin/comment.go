package admin

import (
	"blog/blog_api/service"
	"blog/blog_api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//评论列表
func CommentList(c *gin.Context) {
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
		res.Json(c)
		return
	}
	status, _ := strconv.Atoi(utils.StrVal(json["status"]))
	comment := service.Comment{Status: status}
	count := comment.FindCommentCount()
	result := comment.FindCommentList(json)
	res := &utils.Response{Code: 0, Msg: "", Data: result, Count: count}
	res.Json(c)
}

//审核评论
func CommentReview(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
		res.Json(c)
		return
	}
	comment.UpdateStatus()
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}

//删除评论
func CommentDelete(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
		res.Json(c)
		return
	}
	comment.DeleteComment()
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}
