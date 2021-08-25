package dto

//Login 登录信息
type Login struct {
	UserName string `json:"username" form:"username"` //用户名
	Password string `json:"password" form:"password"` //密码
	Code     string `json:"code" form:"code"`         //验证码
}

//Register 注册信息
type Register struct {
	UserName  string `json:"username" form:"username"`   //用户名
	Password  string `json:"password" form:"password"`   //密码
	Password2 string `json:"password2" form:"password2"` //密码
	Code      string `json:"code" form:"code"`           //验证码
}
