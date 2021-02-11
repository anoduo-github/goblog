package controller

//MessageController 评论控制器
type MessageController struct {
	CommonController
}

//Edit 写留言
func (m *MessageController) Edit() {
	m.TplName = "message.html"
}
