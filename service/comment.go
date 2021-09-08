package service

import (
	"fmt"
	"goblog/dto"
	"goblog/model"
	module "goblog/module/db"
)

//GetParentCommentsByPage 分页获取父评论
func GetParentCommentsByPage(blogId, pageIndex int) ([]model.Comment, error) {
	var comments []model.Comment
	sql := fmt.Sprintf("select * from comments where blog_id=%d and parent_comment_id=-1 order by create_time desc limit %d,%d", blogId, pageIndex*10, 10)
	err := module.GetDB().Raw(sql).Scan(&comments).Error
	return comments, err
}

//GetSonComments 获取子评论
func GetSonComments(blogId, parent_comment_id int) ([]model.Comment, error) {
	var comments []model.Comment
	err := module.GetDB().Where("blog_id=? and parent_comment_id=?", blogId, parent_comment_id).Order("create_time desc").Find(&comments).Error
	return comments, err
}

//CreateCommentDtos 递归构建评论信息
func CreateCommentDtos(parentComments []model.Comment) ([]dto.CommentDto, error) {
	commentDtos := make([]dto.CommentDto, 0)
	for _, v := range parentComments {
		//创建一个commentDto对象
		var commentDto dto.CommentDto
		//将v的值存进去
		commentDto.Comment = v
		//新开一个列表存子评论
		commentDto.Childs = make([]dto.ChildComment, 0)
		//递归获取对应的子评论
		err := recursion(&commentDto.Childs, v.UserName, v.BlogId, v.Id)
		if err != nil {
			return nil, err
		}
		commentDtos = append(commentDtos, commentDto)
	}
	return commentDtos, nil
}

//recursion 递归获取子评论
func recursion(childComments *[]dto.ChildComment, name string, blog_id, parent_comment_id int) error {
	//获取对应的子评论
	childs, err := GetSonComments(blog_id, parent_comment_id)
	if err != nil {
		return err
	}
	for _, v := range childs {
		//赋值给childComment
		var childComment dto.ChildComment
		childComment.Comment = v
		childComment.ParentName = name
		//添加进childComments中
		*childComments = append(*childComments, childComment)
		//递归查找
		err := recursion(childComments, v.UserName, v.BlogId, v.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

//AddComment 添加评论
func AddComment(comment model.Comment) error {
	err := module.GetDB().Create(&comment).Error
	return err
}

//DeleteComment 删除评论
func DeleteComment(id int) error {
	err := module.GetDB().Where("id=?", id).Delete(&model.Comment{}).Error
	return err
}

//CountComment 评论数量
func CountComment(blog_id, parent_comment_id int) (int, error) {
	var counts int
	sql1 := "1=1"
	sql2 := "1=1"
	//表示传的是指定的blog_id，否则就是全查，parent_comment_id同理
	if blog_id != -2 {
		sql1 = fmt.Sprintf("blog_id=%d", blog_id)
	}
	if parent_comment_id != -2 {
		sql2 = fmt.Sprintf("parent_comment_id=%d", parent_comment_id)
	}
	err := module.GetDB().Model(model.Comment{}).Where(sql1 + " and " + sql2).Count(&counts).Error
	return counts, err
}
