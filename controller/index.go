package controller

import (
	"goblog/module/log"
	"goblog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Index 首页
func Index(c *gin.Context) {
	//1.先获取推荐
	recommends, err := service.GetRecommendBlog()
	if err != nil {
		log.Logger.Error("获取最新推荐的博客失败: ", err)
		c.HTML(http.StatusOK, "user/index.html", gin.H{"error": err.Error()})
		return
	}
	//2.在获取博客列表
	index := c.Query("index")
	if index == "" {
		index = "0"
	}
	pageIndex, _ := strconv.Atoi(index)
	blogs, err := service.GetBlog("", "", pageIndex)
	if err != nil {
		log.Logger.Error("获取数据库的博客数据失败: ", err)
		c.HTML(http.StatusOK, "user/index.html", gin.H{"error": err.Error()})
		return
	}
	//查询总条数
	count, err := service.GetAllBlogCount("", "")
	if err != nil {
		log.Logger.Error("获取数据库的博客总条数失败: ", err)
		c.HTML(http.StatusOK, "user/index.html", gin.H{"error": err.Error()})
		return
	}
	//计算总页数
	total := 0
	if count%10 == 0 {
		total = count / 10
		if total == 0 {
			total = 1
		}
	} else {
		total = count/10 + 1
	}

	c.HTML(http.StatusOK, "user/index.html", gin.H{"recommends": recommends, "blogs": blogs, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1})
}

//Admin 管理平台首页
func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{})
}
