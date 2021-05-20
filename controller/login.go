package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//LoginPage 登录页
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

//Login 登录
func Login(c *gin.Context) {

}
