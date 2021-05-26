package service

import (
	"goblog/model"
	module "goblog/module/db"
)

//InsertArticle 插入article
func InsertArticle(a model.Article) error {
	err := module.GetDB().Create(&a).Error
	return err
}

//FindArticleById 根据id查询article
func FindArticleById(id int) (*model.Article, error) {
	var a model.Article
	err := module.GetDB().Where("id=?", id).Take(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}
