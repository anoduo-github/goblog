package service

import (
	"fmt"
	"goblog/model"
	module "goblog/module/db"
)

//GetLinkByPage 获取分页友链
func GetLinkByPage(pageIndex int) ([]model.Link, error) {
	//分页查询
	var links []model.Link
	sql := fmt.Sprintf("select * from links limit %d,%d", pageIndex*10, 10)
	err := module.GetDB().Raw(sql).Scan(&links).Error
	return links, err
}

//GetAllLinkCount 获取友链数量
func GetAllLinkCount() (int, error) {
	var counts int
	err := module.GetDB().Model(model.Link{}).Count(&counts).Error
	return counts, err
}

//GetLinkById 根据id获取友链
func GetLinkById(id int) (model.Link, error) {
	var l model.Link
	err := module.GetDB().Where("id=?", id).Take(&l).Error
	return l, err
}

//AddLink 添加友链
func AddLink(link model.Link) error {
	err := module.GetDB().Create(&link).Error
	return err
}

//UpdateLink 修改友链
func UpdateLink(link model.Link) error {
	err := module.GetDB().Model(&model.Link{}).Where("id=?", link.Id).Updates(&link).Error
	return err
}

//DeleteLink 删除友链
func DeleteLink(id int) error {
	err := module.GetDB().Where("id=?", id).Delete(&model.Link{}).Error
	return err
}

//GetLinkByName 根据name获取友链
func GetLinkByName(name string) (model.Link, error) {
	var l model.Link
	err := module.GetDB().Where("name=?", name).Take(&l).Error
	return l, err
}

//GetAllLink 获取所有友链
func GetAllLink() ([]model.Link, error) {
	var links []model.Link
	err := module.GetDB().Raw("select * from links order by update_time desc").Scan(&links).Error
	return links, err
}
