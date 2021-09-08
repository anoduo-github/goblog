package service

import (
	"errors"
	"fmt"
	"goblog/model"
	module "goblog/module/db"
	"goblog/utils/crypt"
)

//CheckUser 检查用户
func CheckUser(username, password string) (*model.User, error) {
	user, err := FindUserByName(username)
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

//FindUserByName 根据用户名查询
func FindUserByName(username string) (*model.User, error) {
	var user model.User
	err := module.GetDB().Where("user_name=?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUsersByPage 根据模糊用户名,分页获取user(用权限role区分)
func GetUsersByPage(name, role string, pageIndex int) ([]model.User, error) {
	var users []model.User
	sql := fmt.Sprintf("select * from users where role='%s' and user_name like '%s' order by create_date desc limit %d,%d", role, "%"+name+"%", pageIndex*10, 10)
	err := module.GetDB().Raw(sql).Scan(&users).Error
	return users, err
}

//GetUserCountLikeName 根据name模糊查询获取数量(区分权限)
func GetUserCountLikeName(name, role string) (int, error) {
	var counts int
	err := module.GetDB().Model(model.User{}).Where("role=? and user_name like ?", role, "%"+name+"%").Count(&counts).Error
	return counts, err
}

//GetUserById 根据id获取用户
func GetUserById(id int) (model.User, error) {
	var user model.User
	err := module.GetDB().Where("id=?", id).Take(&user).Error
	return user, err
}

//DeleteUser 根据id删除用户
func DeleteUser(id int) error {
	err := module.GetDB().Where("id=?", id).Delete(&model.User{}).Error
	return err
}

//UpdateUser 更新用户
func UpdateUser(user *model.User) error {
	err := module.GetDB().Model(&model.User{}).Where("id=?", user.Id).Updates(user).Error
	return err
}
