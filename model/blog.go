package model

//Blog 文章
type Blog struct {
	Id           int    //主键id
	Author       string //作者
	Title        string //文章标题
	Context      string //文章内容
	Type         string //文章类型
	Tag          string //文章标签
	CreateDate   string //创建文章日期
	UpdateDate   string //修改文章日期
	LikeCount    int    //点赞数
	ViewCount    int    //浏览次数
	Pircture     string //背景图
	Description  string //描述
	CommentCount int    //评论数量
}
