package model

//Like 文章和用户的点赞关系
type Like struct {
	Id      int    //主键
	BlogId  int    //文章id，可重复
	UserIds string //用户id数组字符串，多存点，设置最多存50个
}
