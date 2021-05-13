package controller

/* //AccountController 登录控制器
type AccountController struct {
	CommonController
}

//全局验证码结构体
var cap *captcha.Captcha

//init 函数初始化captcha
func init() {
	//验证码功能
	//使用Beego缓存存储验证码数据
	store := cache.NewMemoryCache()
	//创建验证码
	cap = captcha.NewWithFilter("/captcha/", store)
	//设置验证码长度
	cap.ChallengeNums = 4
	//设置验证码模板高度
	cap.StdHeight = 55
	//设置验证码模板宽度
	cap.StdWidth = 140
}

//LoginPage 登录页
func (a *AccountController) LoginPage() {
	a.TplName = "login.html"
}

//Login 登录
func (a *AccountController) Login() {
	//验证码验证
	if !cap.VerifyReq(a.Ctx.Request) {
		a.putMsg("登录失败，验证码错误", Err)
	}
	//从页面获取信息
	username := strings.TrimSpace(a.GetString("username"))
	password := strings.TrimSpace(a.GetString("password"))
	//从数据库查找数据
	user, err := model.FindByUsername(username)
	if err != nil {
		log.Logger.Errorln("FindByUsername() error:", err)
		a.putMsg("登录失败，查询数据库失败", Err)
	}
	//解密密码
	pwd, err := utils.DePwdCode(user.Password)
	if err != nil {
		log.Logger.Errorln("DePwdCode() error:", err)
		a.putMsg("登录失败，信息解密失败", Err)
	}
	if pwd != password {
		a.putMsg("登录失败，密码错误", Err)
	}
	//加密用户名做cookie
	code, err := utils.EnPwdCode(username)
	if err != nil {
		log.Logger.Errorln("EnPwdCode() error:", err)
		a.putMsg("登录失败，加密信息失败", Err)
	}
	//存一天
	a.Ctx.SetCookie("code", code, 86400)
	//将信息存入redis
	err = redis.SetStruct(code, *user, 3600)
	if err != nil {
		log.Logger.Errorln("SetAccount() error:", err)
		a.putMsg("登录失败，信息保存失败", Err)
	}
	a.putMsg("登录成功", Ok)
}

//RegisterPage 注册页
func (a *AccountController) RegisterPage() {
	a.TplName = "register.html"
}

//Register 注册
func (a *AccountController) Register() {
	//验证码验证
	if !cap.VerifyReq(a.Ctx.Request) {
		a.putMsg("注册失败，验证码错误", Err)
	}
	//从页面获取信息
	username := strings.TrimSpace(a.GetString("username"))
	password1 := strings.TrimSpace(a.GetString("password1"))
	password2 := strings.TrimSpace(a.GetString("password2"))
	//验证格式
	reg1 := regexp.MustCompile(`^[a-zA-Z0-9]{4,8}$`)
	reg2 := regexp.MustCompile(`^[a-zA-Z0-9]{6,10}$`)
	if !reg1.MatchString(username) {
		a.putMsg("注册失败，用户名格式不正确", Err)
	}
	if !reg2.MatchString(password1) {
		a.putMsg("注册失败，密码格式不正确", Err)
	}
	//比较
	if password1 != password2 {
		a.putMsg("注册失败，密码不一致", Err)
	}
	u := &model.User{}
	u.UserName = username
	//查询用户名是否重复
	user, _ := model.FindByUsername(username)
	if user != nil {
		a.putMsg("注册失败，用户名重复", Err)
	}
	//加密
	password, err := utils.EnPwdCode(password1)
	if err != nil {
		log.Logger.Errorln("Register() EnPwdCode() error:", err)
		a.putMsg("注册失败，密码加密失败", Err)
	}
	u.Password = password
	u.UserRole = 0
	u.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	//保存
	err = u.Add()
	if err != nil {
		log.Logger.Errorln("Register() Add() error: %v", err)
		a.putMsg("注册失败，保存数据失败", Err)
	}
	a.putMsg("注册成功", Ok)
}
*/
