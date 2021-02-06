package service

import (
	. "blog/blog_api/db"
	"blog/blog_api/entity"
	"blog/blog_api/utils"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Comment entity.Comment

func (Comment) TableName() string {
	return "Comment"
}

//新增评论
func (comment *Comment) Insert() *gorm.DB {
	return Db.Create(comment)
}

//评论状态审核
func (comment *Comment) UpdateStatus() *gorm.DB {
	return Db.Model(comment).Where("id = ? ", comment.Id).Update("status", 1)
}

//删除评论
func (comment *Comment) DeleteComment() *gorm.DB {
	return Db.Model(comment).Delete(comment)
}


//评论审核list
func (comment *Comment) FindCommentList(Map map[string]interface{}) []Comment {
	comments := make([]Comment, 0)
	page, _ := strconv.Atoi(utils.StrVal(Map["page"]))
	size, _ := strconv.Atoi(utils.StrVal(Map["size"]))
	pageVo := utils.Page{Page: page, Size: size}
	Db.Table("comment c").Select("c.*, b.title as blog_title").
		Joins("left join blog b on c.blog_id = b.id").
		Where("c.status = ?", Map["status"]).
		Order("add_time asc").
		Limit(pageVo.Size).Offset(pageVo.GetStart()).
		Find(&comments)
	return comments
}

//查找评论数量
func (comment *Comment) FindCommentCount() (count int) {
	Db.Model(comment).Where("status = ? ", comment.Status).Count(&count)
	return
}