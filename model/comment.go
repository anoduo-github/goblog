package model

//Comment 评论
type Comment struct {
	Id              int    //主键
	UserName        string //用户名称
	Content         string //内容
	BlogId          int    //博客id,当id为-1表示是留言板评论
	ParentCommentId int    //父评论id,当为-1时表示没有父评论
	Admin           bool   //是否是管理员评论
	CreateTime      string //创建时间
}
