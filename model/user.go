package model

import (
	"github.com/astaxie/beego/orm"
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

//FindByUsername 根据用户名查找指定用户
func FindByUsername(username string) (*User, error) {
	u := new(User)
	err := orm.NewOrm().QueryTable("user").Filter("user_name", username).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Add 添加用户
func (u *User) Add() error {
	_, err := orm.NewOrm().Insert(u)
	return err
}
