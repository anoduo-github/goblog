package service

import (
	"fmt"
	"goblog/model"
	module "goblog/module/db"
)

//GetPictureByPage 获取分页照片
func GetPictureByPage(pageIndex int) ([]model.Picture, error) {
	//分页查询
	var pictures []model.Picture
	sql := fmt.Sprintf("select * from pictures limit %d,%d", pageIndex*10, 10)
	err := module.GetDB().Raw(sql).Scan(&pictures).Error
	return pictures, err
}

//GetAllPictureCount 获取照片数量
func GetAllPictureCount() (int, error) {
	var counts int
	err := module.GetDB().Model(model.Picture{}).Count(&counts).Error
	return counts, err
}

//GetPictureById 根据id获取照片
func GetPictureById(id int) (model.Picture, error) {
	var p model.Picture
	err := module.GetDB().Where("id=?", id).Take(&p).Error
	return p, err
}

//AddPicture 添加照片
func AddPicture(picture model.Picture) error {
	err := module.GetDB().Create(&picture).Error
	return err
}

//UpdatePicture 修改照片
func UpdatePicture(picture model.Picture) error {
	err := module.GetDB().Model(&model.Picture{}).Where("id=?", picture.Id).Updates(&picture).Error
	return err
}

//DeletePicture 删除照片
func DeletePicture(id int) error {
	err := module.GetDB().Where("id=?", id).Delete(&model.Picture{}).Error
	return err
}

//GetPictureByName 根据name获取照片
func GetPictureByName(name string) (model.Picture, error) {
	var p model.Picture
	err := module.GetDB().Where("name=?", name).Take(&p).Error
	return p, err
}

//GetAllPicture 获取所有照片
func GetAllPicture() ([]model.Picture, error) {
	var pictures []model.Picture
	err := module.GetDB().Raw("select * from pictures order by id desc").Scan(&pictures).Error
	return pictures, err
}
