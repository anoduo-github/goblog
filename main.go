package main

import (
	"goblog/module/db"
	"goblog/module/log"
	"goblog/module/redis"
	_ "goblog/router"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//初始化日志
	log.Init()
	//初始化数据库(mysql, redis)
	db.Init()
	redis.Init()
	beego.Run()
}
