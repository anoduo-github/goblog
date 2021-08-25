package service

import (
	"fmt"
	"goblog/model"
	module "goblog/module/db"
)

//Result 结果
type Result struct {
	Type  int
	Total int
}

//AddBlog 添加博客
func AddBlog(a model.Blog) error {
	err := module.GetDB().Create(&a).Error
	return err
}

//FindBlogById 根据id查询blog
func FindBlogById(id int) (model.Blog, error) {
	var a model.Blog
	err := module.GetDB().Where("id=?", id).Take(&a).Error
	return a, err
}

//GetBlog 获取分页博客
func GetBlog(title, typeStr string, pageIndex int) ([]model.Blog, error) {
	//构造标题查询条件
	sql1 := "1=1"
	if title != "" {
		sql1 = fmt.Sprintf("title='%s'", title)
	}
	//构造分类查询条件
	sql2 := "1=1"
	if typeStr != "" {
		sql2 = fmt.Sprintf("type='%s'", typeStr)
	}
	//分页查询
	var blogs []model.Blog
	sql := fmt.Sprintf("select * from blogs where %s and %s order by update_date desc limit %d,%d", sql1, sql2, pageIndex*10, 10)
	err := module.GetDB().Raw(sql).Scan(&blogs).Error
	return blogs, err
}

//GetAllBlogCount 获取博客数量
func GetAllBlogCount(title, typeStr string) (int, error) {
	var counts int
	sql1 := "1=1"
	if title != "" {
		sql1 = fmt.Sprintf("title='%s'", title)
	}
	sql2 := "1=1"
	if typeStr != "" {
		sql2 = fmt.Sprintf("type='%s'", typeStr)
	}
	err := module.GetDB().Model(model.Blog{}).Where(sql1 + " and " + sql2).Count(&counts).Error
	return counts, err
}

//DeleteBlog 删除博客
func DeleteBlog(id int) error {
	err := module.GetDB().Where("id=?", id).Delete(&model.Blog{}).Error
	return err
}

//GetBlogByTitle 根据标题查询博客
func GetBlogByTitle(title string) (model.Blog, error) {
	var blog model.Blog
	err := module.GetDB().Where("title=?", title).Take(&blog).Error
	return blog, err
}

//GetBlogById 根据id查询博客
func GetBlogById(id int) (model.Blog, error) {
	var blog model.Blog
	err := module.GetDB().Where("id=?", id).Take(&blog).Error
	return blog, err
}

//UpdateBlog 更新博客
func UpdateBlog(blog model.Blog) error {
	err := module.GetDB().Model(&model.Blog{}).Where("id=?", blog.Id).Updates(&blog).Error
	return err
}

//GetRecommendBlog 获取推荐的博客
func GetRecommendBlog() ([]model.Blog, error) {
	//根据博客的点赞数，获取四篇点赞数最多的博客
	var blogs []model.Blog
	err := module.GetDB().Order("like_count desc").Limit(4).Find(&blogs).Error
	return blogs, err
}
