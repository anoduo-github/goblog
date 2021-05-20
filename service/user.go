package service

import (
	"errors"
	"goblog/model"
	module "goblog/module/db"
)

//CheckUser 检查用户
func CheckUser(username, password string) (*model.User, error) {
	user, err := Find(username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

//Insert 新增
func Insert(user model.User) error {
	err := module.GetDB().Create(user).Error
	return err
}

//Find 查询
func Find(username string) (*model.User, error) {
	var user model.User
	err := module.GetDB().Where("user_name=?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
