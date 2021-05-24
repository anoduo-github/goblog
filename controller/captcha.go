package controller

import (
	"bytes"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//session session缓存
var session sessions.Session

//initSession 初始化一个session
func initSession(c *gin.Context) {
	if session == nil {
		session = sessions.Default(c)
	}
}

//Captcha 验证码
func Captcha(c *gin.Context) {
	l := captcha.DefaultLen
	w, h := 107, 36
	captchaId := captcha.NewLen(l)
	initSession(c)
	session.Set("captcha", captchaId)
	session.Save()
	serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}

//CaptchaVerify 验证码验证
func CaptchaVerify(code string) bool {
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	}
	return false
}

//serve 服务
func serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
