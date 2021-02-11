package model

//Article 文章
type Article struct {
	Id              int    //主键id
	Author          string //作者
	OrginAuthor     string //原作者
	ArticleTitle    string //文章标题
	ArticleContext  string //文章内容
	ArticleType     string //文章类型
	ArticleTag      string //文章标签
	ArticleCategory string //文章归档
	CreateDate      string //创建文章日期
	UpdateDate      string //修改文章日期
	PreId           int    //上一篇文章id
	NextId          int    //下一篇文章id
	Like            int    //点赞数
}
