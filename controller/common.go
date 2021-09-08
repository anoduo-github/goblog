package controller

import (
	"goblog/module/log"
	"goblog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//返回结果的状态
const (
	statusOk  = 1
	statusErr = 0
)

//FooterInfo 页底信息
func FooterInfo(c *gin.Context) {
	//先获取所有博客
	blogs, err := service.GetAllBlog()
	if err != nil {
		log.Logger.Error("获取所有博客失败: ", err)
		c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
		return
	}

	//遍历计算访问总数，评论总数
	view_count := 0
	comment_count := 0
	for _, v := range blogs {
		view_count = view_count + v.ViewCount
		comment_count = comment_count + v.CommentCount
	}

	//博客总数
	blog_count := len(blogs)

	//获取留言总数
	message_count, err := service.CountComment(-1, -2)
	if err != nil {
		log.Logger.Error("获取所有留言数量失败: ", err)
		c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog_count": blog_count, "view_count": view_count, "comment_count": comment_count, "message_count": message_count, "status": statusOk})
}
