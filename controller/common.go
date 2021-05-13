package controller

/* //返回结果的状态
const (
	Ok  = 0
	Err = -1
)

//CommonController 公共控制器
type CommonController struct {
	beego.Controller
}

//putMsg 提交信息
func (c *CommonController) putMsg(msg interface{}, code int) {
	res := make(map[string]interface{})
	res["code"] = code
	res["msg"] = msg
	c.Data["json"] = res
	c.ServeJSON()
	c.StopRun()
}

//redirect 重定向
func (c *CommonController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

//checkLogin 检查登录
func (c *CommonController) checkLogin() (string, error) {
	//读取cookie中的code
	code := c.Ctx.GetCookie("code")
	if code == "" {
		return "", errors.New("cookie值为空")
	}
	//先从缓存中寻找
	tempname, err := redis.HGetString(code, "UserName")
	if err != nil {
		return "", err
	}
	if tempname != "" {
		return tempname, nil
	}
	//否则从数据库查找
	//解密code
	username, err := utils.DePwdCode(code)
	if err != nil {
		return "", err
	}
	user, err := model.FindByUsername(username)
	if err != nil {
		return "", err
	}
	//存入缓存
	err = redis.SetStruct(code, *user, 3600)
	if err != nil {
		log.Logger.Errorf("SetStruct() error: %v, code: %v\n", err, code)
	}
	return user.UserName, nil
}
*/
