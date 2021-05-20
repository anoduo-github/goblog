package middleware

import (
	"encoding/json"
	"goblog/model"
	"goblog/service"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "one"
var key = "blog20210513"

func GinJwtInit() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "blog jwt",
		Key:         []byte(key),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,

		//PayloadFunc函数：登录期间的回调的函数
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				jsonData, _ := json.Marshal(v)
				return jwt.MapClaims{
					identityKey: string(jsonData),
				}
			}
			return jwt.MapClaims{}
		},

		//IdentityHandler函数：解析并设置用户身份信息
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			var user model.User
			json.Unmarshal([]byte(claims[identityKey].(string)), &user)
			return &user
		},

		//Authenticator函数：根据登录信息对用户进行身份验证的回调函数
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals model.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.UserName
			password := loginVals.Password
			user, err := service.CheckUser(username, password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return user, nil
		},

		//Authorizator函数：接收用户信息并编写授权规则，本项目的API权限控制就是通过该函数编写授权规则的
		Authorizator: func(data interface{}, c *gin.Context) bool {
			/* if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			} */

			return true
		},

		//Unauthorized函数：处理不进行授权的逻辑
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
