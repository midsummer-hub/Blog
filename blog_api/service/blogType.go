package service

import (
	. "blog/blog_api/db"
	"blog/blog_api/entity"
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
