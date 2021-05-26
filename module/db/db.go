package module

import (
	"fmt"
	"goblog/model"
	"goblog/module/config"
	"goblog/module/log"
	"os"

	"github.com/jinzhu/gorm"
)

//_db db连接
var _db *gorm.DB

//Init 初始化mysql连接
func Init() {
	//首先获取配置文件的信息
	//数据库IP
	dbhost := config.Cfg.Section("mysql").Key("host").MustString("localhost")
	//数据库端口
	dbport := config.Cfg.Section("mysql").Key("port").MustString("3306")
	//数据库用户
	dbuser := config.Cfg.Section("mysql").Key("user").MustString("root")
	//数据库密码
	dbpassword := config.Cfg.Section("mysql").Key("password").MustString("root")
	//数据库名
	dbname := config.Cfg.Section("mysql").Key("name").MustString("goblog")
	//设置数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)

	//建立连接
	var err error
	_db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Logger.Errorf("Init() gorm open mysql error [%v]", err)
		os.Exit(1)
	}

	//自动建表
	_db.AutoMigrate(&model.Article{}, &model.User{}, &model.Category{}, &model.Role{}, &model.Tag{})

	_db.Model(&model.Article{}).ModifyColumn("article_context", "text")

	//设置数据库连接池最大连接数
	_db.DB().SetMaxOpenConns(10)
	//设置连接池最大允许的空闲连接数
	_db.DB().SetMaxIdleConns(2)
}

//GetDB 得到一个sqlite连接
func GetDB() *gorm.DB {
	return _db
}
