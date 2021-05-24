package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//返回结果的状态
const (
	statusOk  = 0
	statusErr = -1
)

//Public 微信公众号
func Public(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": statusErr, "msg": "敬请期待"})
}
