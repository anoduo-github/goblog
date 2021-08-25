package middleware

import (
	"fmt"
	"goblog/module/log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//RoleMiddleWare 权限校验中间件
func RoleMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		tokenString := session.Get("token")
		if _, ok := tokenString.(string); ok {
			_, err := GetTokenInfo(tokenString.(string))
			if err != nil {
				log.Logger.Error("get token error: ", err)
				//跳转首页
				c.Redirect(302, "/")
			}
		} else {
			log.Logger.Error("get token error: token is null")
			//跳转首页
			c.Redirect(302, "/")
		}
		fmt.Println("中间件执行了")
		c.Next()
	}
}
