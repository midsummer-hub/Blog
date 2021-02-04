package service

import (
	. "blog/blog_api/db"
	"blog/blog_api/entity"
	"blog/blog_api/utils"
	"github.com/jinzhu/gorm"
)

type BlogType entity.BlogType

func (BlogType) TableName() string {
	return "blog_type"
}


//查找博主
func (blogType *BlogType) FindTypeCount() []map[string]interface{} {
	maps := make([]map[string]interface{},0)
	sql := "select b.`typeId` as typeId,count(`b`.`id`) as b_count,`bt`.`name` `b_name` from `blog` `b` left join `blog_type` `bt` on `b`.`typeId` = `bt`.`id` group by `b`.`typeId`"
	rows,_ := Db.Raw(sql).Rows()
	defer rows.Close()
	for rows.Next() {
		var typeId int
		var bCount int
		var bName string
		rows.Scan(&typeId,&bCount,&bName)
		Map := make(map[string]interface{})
		Map["type_id"] = typeId
		Map["b_count"] = bCount
		Map["b_name"] = bName
		maps = append(maps,Map)
	}
	return maps
}

//查询一个类型
func (blogType *BlogType) FindOne() (bt *BlogType) {
	bt = new(BlogType)
	Db.Where("id = ?",blogType.Id).First(bt)
	return
}

//查询所有类型
func (blogType *BlogType) FindAll() []*BlogType {
	bts := make([]*BlogType, 0)
	Db.Model(blogType).Order("`sort` asc").Find(&bts)
	return bts
}


//类型list
func (blogType *BlogType) FindList(page *utils.Page) []*BlogType {
	bts := make([]*BlogType, 0)
	Db.Model(blogType).Limit(page.Size).Offset(page.GetStart()).Order("`sort` asc").Find(&bts)
	return bts
}


//数量统计
func (blogType *BlogType) Count() (count int) {
	Db.Model(blogType).Count(&count)
	return
}

//添加博客类型
func (blogType *BlogType) Insert() *gorm.DB {
	return Db.Create(blogType)
}

//修改博客类型
func (blogType *BlogType) Update() *gorm.DB {
	return Db.Save(blogType)
}

//删除博客类型
func (blogType *BlogType) Delete() *gorm.DB {
	return Db.Delete(blogType)
}