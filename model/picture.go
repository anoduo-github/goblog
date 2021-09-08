package model

//Picture 照片
type Picture struct {
	Id          int    //唯一id
	Name        string //图片名字
	TimePlace   string //拍摄的时间地点
	Address     string //图片地址
	Description string //图片描述
}
