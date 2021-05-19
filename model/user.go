package model

import (
	"errors"
)

//User 用户
type User struct {
	Id          int    //主键id
	UserName    string //用户名
	Password    string //用户密码
	RecentLogin string //最近登录
	CreateDate  string //创建日期
	UserRole    int    //权限id
}

//Login 登录信息
type Login struct {
	UserName string
	Password string
}

//CheckUser 检查用户
func CheckUser(username, password string) (*User, error) {
	var user User
	err := GetDB().Where("user_name=?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}
