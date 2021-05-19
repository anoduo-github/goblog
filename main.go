package main

import (
	"goblog/module/log"
	"goblog/router"
	"goblog/utils/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	//初始化配置文件
	config.Init("conf/app.ini")
	//初始化日志
	log.Init()
	//初始化数据库(mysql, redis)
	//model.Init()
}

func main() {

	//启动
	router.Run()
}
