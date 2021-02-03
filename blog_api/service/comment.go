package service

import (
	. "blog/blog_api/db"
	"blog/blog_api/entity"
	"github.com/jinzhu/gorm"
)

type Comment entity.Comment

func (Comment) TableName() string {
	return "Comment"
}

//新增评论
func (comment *Comment) Insert() *gorm.DB {
	return Db.Create(comment)

}