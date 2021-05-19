package controller

import (
	"bytes"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//LoginPage 登录页
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login2.html", gin.H{})
}

//Login 登录
func Login(c *gin.Context) {

}

//Captcha 验证码
func Captcha(c *gin.Context) {
	l := captcha.DefaultLen
	w, h := 107, 36
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	session.Save()
	serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}

//CaptchaVerify 验证码验证
func CaptchaVerify(c *gin.Context) {
	code := c.Query("code")
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": -1, "msg": "failed"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"status": -1, "msg": "failed"})
	}
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
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
