package service

import (
	"goblog/model"
	module "goblog/module/db"
	"strconv"
	"strings"
)

//GetLikes 获取点赞关系（根据文章id）
func GetLikes(blog_id int) ([]model.Like, error) {
	var likes []model.Like
	err := module.GetDB().Model(model.Like{}).Where("blog_id=?", blog_id).Order("id desc").Scan(&likes).Error
	return likes, err
}

//UpdateLike 更新like
func UpdateLike(like model.Like) error {
	err := module.GetDB().Model(&model.Like{}).Where("id=?", like.Id).Updates(&like).Error
	return err
}

//AddLike 添加like
func AddLike(blog_id int, user_id string) error {
	var like model.Like
	like.BlogId = blog_id
	like.UserIds = user_id
	//存入
	err := module.GetDB().Create(&like).Error
	return err
}

//CheckLike 检查用户是否点赞过某文章
func CheckLike(blog_id, user_id int) (bool, error) {
	likes, err := GetLikes(blog_id)
	if err != nil {
		return false, err
	}
	userIdStr := strconv.Itoa(user_id)
	//遍历likes判断UserIds是否存在指定的user_id
	for _, v := range likes {
		arrs := strings.Split(v.UserIds, ",")
		for _, idStr := range arrs {
			if idStr == userIdStr {
				return true, nil
			}
		}
	}
	return false, nil
}
