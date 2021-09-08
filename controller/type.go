package controller

import (
	"goblog/module/log"
	"goblog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Types 分类管理列表页
func Types(c *gin.Context) {
	//获取参数
	index := c.Param("index")
	if index == "" {
		index = "0"
	}
	pageIndex, _ := strconv.Atoi(index)

	//获取分页types
	types, err := service.GetTypeByPage(pageIndex)
	if err != nil {
		log.Logger.Error("获取所有分类失败: ", err)
		c.HTML(http.StatusOK, "admin/types.html", gin.H{"error": err.Error()})
		return
	}

	//获取类型总数
	count, err := service.GetAllTypeCount()
	if err != nil {
		log.Logger.Error("获取分类总数失败: ", err)
		c.HTML(http.StatusOK, "admin/types.html", gin.H{"error": err.Error()})
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

	c.HTML(http.StatusOK, "admin/types.html", gin.H{"types": types, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1})
}

//TypeAdd 添加分类
func TypeAdd(c *gin.Context) {
	//获取参数
	name := c.Param("name")
	//查重
	t, err := service.GetTypeByName(name)
	if err != nil && err.Error() != "record not found" {
		log.Logger.Error("获取类型失败: ", err)
		c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
		return
	}
	if t.Name != "" {
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "类型名称重复"})
		return
	}

	//保存type
	err = service.AddType(name)
	if err != nil {
		log.Logger.Error("add type error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": statusOk, "msg": "新增成功"})
}

//TypeEdit 编辑分类
func TypeEdit(c *gin.Context) {
	//获取参数
	name := c.Param("name")
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.UpdateType(id, name)
	if err != nil {
		log.Logger.Error("edit type error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": statusOk, "msg": "修改成功"})
}

//EditPage 编辑页面
func EditPage(c *gin.Context) {
	idStr := c.Param("id")
	c.HTML(http.StatusOK, "admin/types-input.html", gin.H{"id": idStr})
}

//DeleteType 删除type
func DeleteType(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.DeleteType(id)
	if err != nil {
		log.Logger.Error("delete type error: ", err)
		return
	}
	c.Redirect(302, "/type")
}
