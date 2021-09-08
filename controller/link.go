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

//Links 友链管理列表页
func Links(c *gin.Context) {
	//获取参数
	index := c.Param("index")
	if index == "" {
		index = "0"
	}
	pageIndex, _ := strconv.Atoi(index)

	//获取分页友链
	links, err := service.GetLinkByPage(pageIndex)
	if err != nil {
		log.Logger.Error("获取所有友链失败: ", err)
		c.HTML(http.StatusOK, "admin/friendlinks.html", gin.H{"error": err.Error()})
		return
	}

	//获取友链总数
	count, err := service.GetAllLinkCount()
	if err != nil {
		log.Logger.Error("获取友链总数失败: ", err)
		c.HTML(http.StatusOK, "admin/friendlinks.html", gin.H{"error": err.Error()})
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

	c.HTML(http.StatusOK, "admin/friendlinks.html", gin.H{"links": links, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1})
}

//EditLink 编辑友链
func EditLink(c *gin.Context) {
	//获取id
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	//跳到新增
	if id == 0 {
		c.HTML(http.StatusOK, "admin/friendlinks-input.html", gin.H{})
		return
	}

	//根据id获取link
	link, err := service.GetLinkById(id)
	if err != nil {
		log.Logger.Error("获取友链失败: ", err)
		c.HTML(http.StatusOK, "admin/friendlinks-input.html", gin.H{"error": err.Error()})
		return
	}

	//跳到修改
	c.HTML(http.StatusOK, "admin/friendlinks-input.html", gin.H{"link": link})
}

//AddLink 添加或修改友链
func AddLink(c *gin.Context) {
	//获取表单数据
	idStr := c.PostForm("id")
	name := c.PostForm("name")
	address := c.PostForm("address")
	img := c.PostForm("img")

	//创建link
	var link model.Link
	link.Address = address
	link.Name = name
	link.Img = img

	//判断idStr,为空表示新增，否则为修改
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		link.Id = id
		link.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
		err := service.UpdateLink(link)
		if err != nil {
			log.Logger.Error("更新友链失败: ", err)
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
	} else {
		//查重
		l, err := service.GetLinkByName(name)
		if err != nil && err.Error() != "record not found" {
			log.Logger.Error("获取友链失败: ", err)
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
		if l.Name != "" {
			c.JSON(http.StatusOK, gin.H{"msg": "博客名称重复", "status": statusErr})
			return
		}

		link.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		link.UpdateTime = link.CreateTime
		err = service.AddLink(link)
		if err != nil {
			log.Logger.Error("新增友链失败: ", err)
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"msg": "提交成功", "status": statusOk})
}

//DeleteLink 删除友链
func DeleteLink(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.DeleteLink(id)
	if err != nil {
		log.Logger.Error("删除友链失败: ", err)
	}
	c.Redirect(302, "/link/list/0")
}

//LinkPage 友链页面
func LinkPage(c *gin.Context) {
	links, err := service.GetAllLink()
	if err != nil {
		log.Logger.Error("获取所有友链失败: ", err)
		c.HTML(http.StatusOK, "user/friends.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "user/friends.html", gin.H{"links": links})
}
