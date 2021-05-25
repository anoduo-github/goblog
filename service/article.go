package service

import (
	"goblog/model"
	module "goblog/module/db"
)

//InsertArticle 插入
func InsertArticle(a model.Article) error {
	err := module.GetDB().Create(&a).Error
	return err
}
