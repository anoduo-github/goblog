package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//Session session中间件
func Session(keyPairs string) gin.HandlerFunc {
	store := cookie.NewStore([]byte("very good"))
	store.Options(sessions.Options{
		MaxAge: 2 * 3600, //seconds
		Path:   "/",
	})
	return sessions.Sessions(keyPairs, store)
}
