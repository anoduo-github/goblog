package service

import (
	"fmt"
	"goblog/model"
	module "goblog/module/db"

	"github.com/jinzhu/gorm"
)

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

//GetAllBlog 获取所有博客
func GetAllBlog() ([]model.Blog, error) {
	var blogs []model.Blog
	err := module.GetDB().Raw("select * from blogs").Scan(&blogs).Error
	return blogs, err
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

//GetBlogLikeTitle 根据标题进行模糊查询,分页查询
func GetBlogLikeTitle(title string, pageIndex int) ([]model.Blog, error) {
	var blogs []model.Blog
	sql := fmt.Sprintf("select * from blogs where title like '%s' order by update_date desc limit %d,%d", "%"+title+"%", pageIndex*10, 10)
	err := module.GetDB().Raw(sql).Scan(&blogs).Error
	return blogs, err
}

//CountBlogLikeTitle 计算模糊查询title的blog数量
func CountBlogLikeTitle(title string) (int, error) {
	var counts int
	err := module.GetDB().Model(model.Blog{}).Where("title like ?", "%"+title+"%").Count(&counts).Error
	return counts, err
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

//GetArchivesBlog 获取时间轴需要的博客
func GetArchivesBlog() ([]model.Blog, error) {
	//按更新时间升序排列，取五个
	var blogs []model.Blog
	sql := "select * from blogs order by update_date asc limit 5"
	err := module.GetDB().Raw(sql).Scan(&blogs).Error
	return blogs, err
}

//UpdateBlogLikeCount 更新博客点赞数+1
func UpdateBlogLikeCount(blog_id int) error {
	err := module.GetDB().Model(&model.Blog{}).Where("id=?", blog_id).Update("like_count", gorm.Expr("like_count + 1")).Error
	return err
}

//UpdateBlogCommentCount 更新博客评论数+1
func UpdateBlogCommentCount(blog_id int) error {
	err := module.GetDB().Model(&model.Blog{}).Where("id=?", blog_id).Update("comment_count", gorm.Expr("comment_count + 1")).Error
	return err
}
