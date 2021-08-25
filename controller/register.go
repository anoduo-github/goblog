package controller

import (
	"goblog/dto"
	"goblog/model"
	"goblog/module/log"
	"goblog/service"
	"goblog/utils/crypt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//RegisterPage 注册页
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.html", gin.H{})
}

//Register 注册
func Register(c *gin.Context) {
	r := dto.Register{}
	//绑定参数
	err := c.ShouldBind(&r)
	if err != nil {
		log.Logger.Error("bind struct error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "注册失败，获取表单数据失败"})
		return
	}

	//验证码
	if !CaptchaVerify(r.Code) {
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "注册失败，验证码不对"})
		return
	}

	//比较两次的密码
	if r.Password != r.Password2 {
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "注册失败，密码不一致"})
		return
	}

	//查查
	u, _ := service.FindUser(r.UserName)
	if u != nil {
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "注册失败，用户名重复"})
		return
	}

	//保存注册信息
	user := model.User{}
	user.UserName = r.UserName
	user.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	user.Role = "user"
	//加密
	pwd, err := crypt.EnPwdCode(r.Password)
	if err != nil {
		log.Logger.Error("EnPwdCode error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "注册失败，数据处理失败"})
		return
	}
	user.Password = pwd
	err = service.InsertUser(user)
	if err != nil {
		log.Logger.Error("Insert sql error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "注册失败，注册信息保存失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": statusOk, "msg": "注册成功"})
}
