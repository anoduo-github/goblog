package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//返回结果的状态
const (
	statusOk  = 1
	statusErr = 0
)

//博客类型
const (
	GO    = 0
	MYSQL = 1
	LINUX = 2
	HTTP  = 3
	OTHER = 4
)

//Public 微信公众号
func Public(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "敬请期待"})
}
