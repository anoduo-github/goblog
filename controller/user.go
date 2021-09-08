package controller

import (
	"goblog/module/log"
	"goblog/service"
	"goblog/utils/crypt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Users 用户列表(查询的是普通用户)
func Users(c *gin.Context) {
	//首先获取url参数
	index := c.DefaultQuery("index", "0")
	name := c.DefaultQuery("name", "")
	pageIndex, _ := strconv.Atoi(index)

	//查询
	users, err := service.GetUsersByPage(name, "user", pageIndex)
	if err != nil {
		log.Logger.Error("获取数据库的用户数据失败: ", err)
		c.HTML(http.StatusOK, "admin/users.html", gin.H{"error": err.Error()})
		return
	}
	count, err := service.GetUserCountLikeName(name, "user")
	if err != nil {
		log.Logger.Error("获取数据库的用户总条数失败: ", err)
		c.HTML(http.StatusOK, "admin/users.html", gin.H{"error": err.Error()})
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

	c.HTML(http.StatusOK, "admin/users.html", gin.H{"users": users, "name": name, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1})
}

//DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.DeleteUser(id)
	if err != nil {
		log.Logger.Error("删除用户失败: ", err)
	}
	c.Redirect(302, "/user/list")
}

//UserDetail 用户详情
func UserDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, err := service.GetUserById(id)
	if err != nil {
		log.Logger.Error("获取数据库的用户信息失败: ", err)
		c.HTML(http.StatusOK, "admin/users-detail.html", gin.H{"error": err.Error()})
		return
	}
	//解密用户密码
	user.Password, err = crypt.DePwdCode(user.Password)
	if err != nil {
		log.Logger.Error("解密用户密码失败: ", err)
		c.HTML(http.StatusOK, "admin/users-detail.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "admin/users-detail.html", gin.H{"user": user})
}

//About 关于我
func About(c *gin.Context) {
	c.HTML(http.StatusOK, "user/about.html", gin.H{})
}

//UserInfo 用户信息
func UserInfo(c *gin.Context) {
	//判断用户是否登录，显示登录信息
	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusOK, gin.H{"msg": "未登录", "status": statusErr})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": username, "status": statusOk})
	}
}
