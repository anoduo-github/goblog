package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Error 转到错误
func Error(c *gin.Context) {
	name := c.Param("name")
	errMsg := ""
	switch name {
	case "时间轴", "留言板", "友链", "照片墙", "关于我":
		errMsg = name + "：功能还未实现，请等待"
	case "role":
		errMsg = "用户权限不足，请登录管理员账号访问"
	case "check":
		errMsg = "检测不到用户存在，请先<a href='/login'>登录</a>"
	default:
		errMsg = "异常错误，请联系管理员查看项目日志"
	}
	c.HTML(http.StatusOK, "error/error.html", gin.H{"error": errMsg})
}
