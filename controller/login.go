package controller

import (
	"goblog/dto"
	"goblog/module/log"
	"goblog/service"
	"goblog/utils/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//LoginPage 登录页
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.html", gin.H{})
}

//Login 登录
func Login(c *gin.Context) {
	l := dto.Login{}

	//绑定参数
	err := c.ShouldBind(&l)
	if err != nil {
		log.Logger.Error("bind struct error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "登录失败，获取表单数据失败"})
		return
	}

	//验证码
	if !CaptchaVerify(l.Code) {
		log.Logger.Error("验证码校验失败: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "登录失败，验证码不对"})
		return
	}

	//校验用户
	u, err := service.CheckUser(l.UserName, l.Password)
	if err != nil {
		log.Logger.Error("校验用户失败: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": err.Error()})
		return
	}

	//生成token令牌
	tokenString, err := middleware.SaveToken(u.Role)
	if err != nil {
		log.Logger.Error("生产token令牌失败: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": err.Error()})
		return
	}

	//更新用户登录时间
	u.RecentLogin = time.Now().Format("2006-01-02 15:04:05")
	err = service.UpdateUser(u)
	if err != nil {
		log.Logger.Error("数据库更新用户信息失败: ", err)
	}

	//存入session
	session := sessions.Default(c)
	session.Set("token", tokenString)
	session.Set("username", u.UserName)
	session.Set("user_id", u.Id)
	session.Set("recent", u.RecentLogin)
	err = session.Save()
	if err != nil {
		log.Logger.Error("session保存失败: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": statusOk, "msg": "登录成功"})
}

//Logout 注销
func Logout(c *gin.Context) {
	//获取session
	session := sessions.Default(c)
	//设置立刻过期
	session.Clear()
	session.Save()
	c.Redirect(302, "/")
}
