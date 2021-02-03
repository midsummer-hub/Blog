package utils

import "math"

type Page struct {
	//当前页
	Page int `json:"page"`
	//每页的记录数
	Size int `json:"size"`
	//总记录数
	Total int
}

func (page *Page) GetPage() int {
	//最大页数
	max := int(math.Ceil(float64(page.Total)/float64(page.Size)))
	//当前陎大于最大页数 取最大值
	if page.Page > max {
		page.Page = max
	}
	return page.Page
}

//读取数据起始位置
func (page *Page) GetStart() int {
	return (page.Page - 1) * page.Size
}



