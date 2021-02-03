package entity

type Comment struct {
	Id        int    `gorm:"column:id"json:"id"`
	Ip        string `gorm:"column:ip"json:"ip"`
	Content   string `gorm:"column:content"json:"content"`
	BlogId    int    `gorm:"column:blog_id"json:"blog_id"`
	Status    int    `gorm:"column:status"json:"status"`
	AddTime   string `gorm:"column:add_time"json:"add_time"`
	BlogTitle string `gorm:"-"json:"blog_title"`
}
