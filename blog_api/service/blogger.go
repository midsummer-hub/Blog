package service

import (
	. "blog/blog_api/db"
	"blog/blog_api/entity"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Blogger entity.Blogger

func (Blogger) TableName() string {
	return "blogger"
}

//登录
func (blogger *Blogger) Login() (b *Blogger) {
	b = new(Blogger)
	fmt.Println(b)
	Db.Where("username = ?",blogger.Username).First(b)
	return
}


//查找博主
func (blogger *Blogger) Find() (b *Blogger) {
	b = new(Blogger)
	fmt.Println(b)
	Db.Where("id = 1").First(b)
	return
}

//插入博主信息
func (blogger *Blogger) Insert() *gorm.DB {
	return Db.Create(blogger)
}

//修改用户信息
func (blogger *Blogger) UpdateInfo() *gorm.DB {
	if blogger.Password != ""{
		return Db.Model(blogger).Update(blogger)
	}
	return Db.Save(blogger)
}

//修改用户密码
func (blogger *Blogger) UpdatePassword() *gorm.DB {
	if blogger.Password != ""{
		return Db.Model(blogger).Select("password").Update(blogger)
	}
	return Db.Save(blogger)
}