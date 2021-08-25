package model

//Comment 评论
type Comment struct {
	Id              int    //主键
	UserId          int    //用户id
	Content         string //内容
	BlogId          int    //博客id
	ParentCommentId int    //父评论id
	AdminCommentId  int    //管理员评论id
	CreateTime      string //创建时间
}
