package model

//User 用户
type User struct {
	Id          int    //主键id
	UserName    string //用户名
	Password    string //用户密码
	RecentLogin string //最近登录
	CreateDate  string //创建日期
	UserRole    int    //权限id
}
