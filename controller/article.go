package controller

import "goblog/module/log"

//ArticleController 博客页控制器
type ArticleController struct {
	CommonController
}

//List 博客列表
func (a *ArticleController) List() {
	//检查登录情况
	username, err := a.checkLogin()
	if err != nil {
		log.Logger.Errorln("checkLogin() error:", err)
	}
	a.Data["username"] = username
	a.TplName = "article.html"
}

//Details 博客详情页
func (a *ArticleController) Details() {
	a.TplName = "read.html"
}
