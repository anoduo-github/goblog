package model

//Link 友链
type Link struct {
	Id         int    //唯一id
	Name       string //博客名称
	Address    string //博客地址
	Img        string //博客的背景图片的地址
	CreateTime string //创建时间
	UpdateTime string //修改时间
}
