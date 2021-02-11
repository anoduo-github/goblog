package db

import (
	"fmt"
	"goblog/model"
	"goblog/module/log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Init 初始化mysql数据库连接
func Init() {
	//首先获取配置文件的信息
	//数据库IP
	dbhost := beego.AppConfig.String("db.host")
	//数据库端口
	dbport := beego.AppConfig.String("db.port")
	//端口检查,为空则默认3306
	if dbport == "" {
		dbport = "3306"
	}
	//数据库用户
	dbuser := beego.AppConfig.String("db.user")
	//数据库密码
	dbpassword := beego.AppConfig.String("db.password")
	//数据库名
	dbname := beego.AppConfig.String("db.name")
	//设置数据库连接
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", dbuser, dbpassword, dbhost, dbport, dbname)
	//注册默认数据库，驱动为mysql, 第三个参数就是我们的数据库连接字符串。
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		log.Logger.Errorln("orm.RegisterDataBase() is error:", err)
		os.Exit(0)
	}
	//自动建表
	orm.RegisterModel(new(model.Article))
	orm.RegisterModel(new(model.User))
	orm.RegisterModel(new(model.Category))
	orm.RegisterModel(new(model.Role))
	orm.RegisterModel(new(model.Tag))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Logger.Errorln("orm.RunSyncdb() is error:", err)
		os.Exit(0)
	}
	//设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 20)
	//设置数据库的最大数据库连接
	orm.SetMaxOpenConns("default", 100)
	// 默认运行模式
	if beego.AppConfig.String("db.debug") == "true" {
		orm.Debug = true
	}
}
