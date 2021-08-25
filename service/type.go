package service

import (
	"errors"
	"goblog/model"
	module "goblog/module/db"
)

//GetAllType 获取所有类型
func GetAllType() ([]model.Type, error) {
	types := make([]model.Type, 0)
	err := module.GetDB().Find(&types).Error
	return types, err
}

//GetTypeById 根据id获取type
func GetTypeById(id int) (model.Type, error) {
	var t model.Type
	err := module.GetDB().Where("id=?", id).Take(&t).Error
	return t, err
}

//GetTypeByName 根据name获取type
func GetTypeByName(name string) (model.Type, error) {
	var t model.Type
	err := module.GetDB().Where("name=?", name).Take(&t).Error
	return t, err
}

//AddType 添加type
func AddType(name string) error {
	t := model.Type{
		Name: name,
	}
	temp, _ := GetTypeByName(name)
	if temp.Name == name {
		return errors.New("分类名已存在")
	}
	err := module.GetDB().Create(&t).Error
	return err
}

//UpdateType 修改type
func UpdateType(id int, name string) error {
	err := module.GetDB().Model(&model.Type{}).Where("id=?", id).Update("name", name).Error
	return err
}

//DeleteType 删除type
func DeleteType(id int) error {
	err := module.GetDB().Where("id=?", id).Delete(&model.Type{}).Error
	return err
}
