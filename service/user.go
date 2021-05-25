package service

import (
	"errors"
	"goblog/model"
	module "goblog/module/db"
	"goblog/utils/crypt"
)

//CheckUser 检查用户
func CheckUser(username, password string) (*model.User, error) {
	user, err := FindUser(username)
	if err != nil {
		return nil, err
	}
	//密码解密
	pwd, err := crypt.DePwdCode(user.Password)
	if err != nil {
		return nil, err
	}
	if pwd != password {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

//InsertUser 新增
func InsertUser(user model.User) error {
	err := module.GetDB().Create(&user).Error
	return err
}

//FindUser 查询
func FindUser(username string) (*model.User, error) {
	var user model.User
	err := module.GetDB().Where("user_name=?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
