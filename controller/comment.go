package controller

import (
	"goblog/model"
	"goblog/module/log"
	"goblog/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//Comments 留言板
func Comments(c *gin.Context) {
	index := c.Param("index")
	pageIndex, _ := strconv.Atoi(index)
	//首先获取10条父评论（一页的标准）
	parentComments, err := service.GetParentCommentsByPage(-1, pageIndex)
	if err != nil {
		log.Logger.Error("获取留言板信息失败: ", err)
		c.HTML(http.StatusOK, "user/message.html", gin.H{"error": err})
		return
	}
	//获取构造好的结果
	commentDtos, err := service.CreateCommentDtos(parentComments)
	if err != nil {
		log.Logger.Error("构造留言板信息失败: ", err)
		c.HTML(http.StatusOK, "user/message.html", gin.H{"error": err})
		return
	}
	//获取用户名
	username := c.GetString("username")
	//获取所有评论数量
	count, err := service.CountComment(-1, -2)
	if err != nil {
		log.Logger.Error("获取留言板信息总数失败: ", err)
		c.HTML(http.StatusOK, "user/message.html", gin.H{"error": err})
		return
	}
	//获取所有父评论数量
	parent_count, err := service.CountComment(-1, -1)
	if err != nil {
		log.Logger.Error("获取留言板父评论信息总数失败: ", err)
		c.HTML(http.StatusOK, "user/message.html", gin.H{"error": err})
		return
	}
	//计算页数
	//因为分页标准是按照每页父评论数量10条规定的
	//所以用所有父评论的数量来计算页数
	total := 0
	if parent_count%10 == 0 {
		total = parent_count / 10
		//防止显示总页数为0，因为不管有没有评论，总要展示第一页，要是总页数为0，岂不是显示成1/0,应该是1/1
		if total == 0 {
			total = 1
		}
	} else {
		total = parent_count/10 + 1
	}

	c.HTML(http.StatusOK, "user/message.html", gin.H{"comments": commentDtos, "username": username, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1})
}

//AddComment 添加评论
func AddComment(c *gin.Context) {
	//获取参数
	content := c.PostForm("content")
	blog_id_str := c.DefaultPostForm("blog_id", "-1")
	blog_id, _ := strconv.Atoi(blog_id_str)
	parent_comment_id_str := c.DefaultPostForm("parent_id", "-1")
	parent_comment_id, _ := strconv.Atoi(parent_comment_id_str)

	//构建评论信息
	var comment model.Comment
	comment.Content = content
	comment.BlogId = blog_id
	comment.ParentCommentId = parent_comment_id

	//获取用户名
	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusOK, gin.H{"msg": "请先<a href='/login'>登录</a>", "status": statusErr})
		return
	}
	if username == "admin" {
		comment.Admin = true
	} else {
		comment.Admin = false
	}
	comment.UserName = username
	comment.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	//插入
	err := service.AddComment(comment)
	if err != nil {
		log.Logger.Error("插入留言板信息失败: ", err)
		c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
		return
	}
	//更新博客的评论数量
	if blog_id != -1 {
		err := service.UpdateBlogCommentCount(blog_id)
		if err != nil {
			log.Logger.Errorf("blog_id=%d,更新博客评论数失败: %s", blog_id, err.Error())
		}
	}
	c.JSON(http.StatusOK, gin.H{"msg": "发布成功", "status": statusOk})
}

//DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.DeleteComment(id)
	if err != nil {
		log.Logger.Error("删除留言板信息失败: ", err)
	}
	c.Redirect(302, "/comment/0")
}
