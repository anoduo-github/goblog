package service

import (
	"errors"
	"goblog/model"
	module "goblog/module/db"
)

//CheckUser 检查用户
func CheckUser(username, password string) (*model.User, error) {
	var user model.User
	err := module.GetDB().Where("user_name=?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}
