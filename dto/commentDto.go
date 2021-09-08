package dto

import "goblog/model"

//ChildComment 子评论信息
type ChildComment struct {
	model.Comment
	ParentName string //父评论用户名
}

//CommentDto 评论信息
type CommentDto struct {
	model.Comment
	Childs []ChildComment //子评论
}
