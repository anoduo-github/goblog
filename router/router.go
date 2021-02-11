package router

import (
	"goblog/controller"

	"github.com/astaxie/beego"
)

//init 初始化
func init() {
	//首页
	beego.Router("/", &controller.IndexController{}, "Get:Index")
	beego.Router("/index", &controller.IndexController{}, "Get:Index")

	//登录
	beego.AutoRouter(&controller.AccountController{})
	//博客
	beego.AutoRouter(&controller.ArticleController{})
	//日记
	beego.AutoRouter(&controller.DiaryController{})
	//留言
	beego.AutoRouter(&controller.MessageController{})
}
