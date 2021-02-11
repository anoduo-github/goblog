package controller

//IndexController 首页控制器
type IndexController struct {
	CommonController
}

//Index 首页
func (index *IndexController) Index() {
	index.TplName = "index.html"
}
