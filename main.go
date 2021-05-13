package main

import (
	"goblog/module/db"
	"goblog/module/log"
	_ "goblog/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//初始化日志
	log.Init()
	//初始化数据库(mysql, redis)
	db.Init()
}
