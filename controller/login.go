package controller

import (
	"goblog/dto"
	"goblog/module/log"
	"goblog/service"
	"goblog/utils/middleware"
	"net/http"

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
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "登录失败，验证码不对"})
		return
	}

	//校验用户
	u, err := service.CheckUser(l.UserName, l.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": err.Error()})
		return
	}

	//生成token令牌
	tokenString, err := middleware.SaveToken(u.Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": err.Error()})
		return
	}

	//存入session
	session := sessions.Default(c)
	session.Set("token", tokenString)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": statusOk, "msg": "登录成功"})
}
