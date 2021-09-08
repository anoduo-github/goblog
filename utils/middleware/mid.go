package middleware

import (
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
			role, err := GetTokenInfo(tokenString.(string))
			if err != nil {
				log.Logger.Error("获取token失败: ", err)
				//跳转错误页
				c.Redirect(302, "/error/err")
			}
			if role != "admin" {
				//跳转错误页
				c.Redirect(302, "/error/role")
			}
		} else {
			//表示未登录，跳转到登录页
			c.Redirect(302, "/login")
		}

		//判定成功后
		c.Next()
	}
}

//UserMiddleWare 用户检测中间件
func UserMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if v1, ok := username.(string); ok {
			c.Set("username", v1)
		}
		user_id := session.Get("user_id")
		if v2, ok := user_id.(int); ok {
			c.Set("user_id", v2)
		}
		recent := session.Get("recent")
		if v3, ok := recent.(string); ok {
			c.Set("recent", v3)
		}
		c.Next()
	}
}
