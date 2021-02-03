package db

import (
	"blog/blog_api/conf"
	"blog/blog_api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init()  {
	var err error
	var dbConfig = conf.Conf.Db
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)
	Db,err = gorm.Open(conf.Conf.Db.Dialects,url)
	if err != nil {
		panic(err)
	}
	if Db.Error != nil {
		panic(Db.Error)
	}

	Db.DB().SetMaxIdleConns(dbConfig.MaxIdle)		//最大空闲数
	Db.DB().SetMaxOpenConns(dbConfig.MaxOpen)		//最大连接数
	logger := utils.Log()
	Db.SetLogger(logger)
	Db.LogMode(true)
	logger.Info("mysql connect success")

}