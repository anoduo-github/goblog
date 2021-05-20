package main

import (
	"goblog/module/config"
	module "goblog/module/db"
	"goblog/module/log"
	"goblog/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	//初始化配置文件
	config.Init("conf/app.ini")
	//初始化日志
	log.Init()
	//初始化数据库(mysql, redis)
	module.Init()
}

func main() {
	//启动
	router.Run()
}
