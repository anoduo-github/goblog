package controller

import (
	"goblog/model"
	"goblog/module/log"
	"goblog/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddLike(c *gin.Context) {
	//获取参数
	blogIdStr := c.PostForm("blogId")
	userIdStr := c.PostForm("userId")
	blogId, _ := strconv.Atoi(blogIdStr)

	//获取指定博客id的like列表（按id降序排列的）
	likes, err := service.GetLikes(blogId)

	if err != nil {
		log.Logger.Error("获取likes失败: ", err)
		c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
		return
	} else if len(likes) == 0 {
		//新增
		err := service.AddLike(blogId, userIdStr)
		if err != nil {
			log.Logger.Error("添加likes失败: ", err)
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
	} else {
		arrs := strings.Split(likes[0].UserIds, ",")
		if len(arrs) == 50 {
			//新增
			err := service.AddLike(blogId, userIdStr)
			if err != nil {
				log.Logger.Error("添加likes失败: ", err)
				c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
				return
			}
		} else {
			//修改
			var like model.Like
			like.Id = likes[0].Id
			like.BlogId = blogId
			like.UserIds = likes[0].UserIds + "," + userIdStr
			err := service.UpdateLike(like)
			if err != nil {
				log.Logger.Error("更新likes失败: ", err)
				c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
				return
			}
		}
	}

	//更新博客点赞数
	err = service.UpdateBlogLikeCount(blogId)
	if err != nil {
		log.Logger.Errorf("blogId=%d,更新博客点赞数失败: %s", blogId, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"msg": "提交成功", "status": statusOk})
}
