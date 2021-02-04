package entity

type BlogType struct {
	Id   int    `gorm:"column:id"json:"id"`
	Name string `gorm:"column:name"json:"name"`
	Sort int    `gorm:"column:sort"json:"sort"`
}
