package router

import (
	"fmt"
	"goblog/controller"
	"goblog/module/config"
	"goblog/module/log"
	"goblog/utils/middleware"
	"goblog/utils/session"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	//设置debug模式
	gin.SetMode(gin.DebugMode)

	//创建路由
	r := gin.Default()

	//注册中间件
	r.Use(session.Session("Lee"))

	//设置静态资源文件目录，并且绑定一个Url前缀
	r.Static("static", "static")

	//加载templates目录下面的所有模版文件，模版文件扩展名随意
	r.LoadHTMLGlob("templates/**/*")

	//初始化路由绑定
	initRouter(r)

	//监听端口，启动服务
	port := config.Cfg.Section("blog").Key("port").MustInt(8080)
	err := r.Run(fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Logger.Errorf("gin could not Run error: %v", err)
		os.Exit(1)
	}
}

func initRouter(r *gin.Engine) {
	//首页
	r.GET("/", controller.Index)

	//登录页
	r.GET("/login", controller.LoginPage)
	//登录
	r.POST("/dologin", controller.Login)
	//验证码
	r.GET("/captcha", controller.Captcha)

	//注册页
	r.GET("/register", controller.RegisterPage)
	//注册
	r.POST("/doregister", controller.Register)

	//博客列表页
	r.GET("/blog/list", controller.BlogList)
	//提交博客
	r.POST("/blog/submit", controller.BlogSubmit)
	//编辑博客
	r.GET("/blog/edit", controller.BlogEdit)
	r.GET("/blog/edit/:id", controller.BlogEdit)
	//博客详情
	r.GET("/blog/details/:id", controller.BlogDetails)
	//上传图片
	r.POST("/blog/upload/picture", controller.UploadPicture)
	//删除博客
	r.GET("blog/delete/:id", controller.DeleteBlog)

	//微信公众号链接
	r.GET("/public", controller.Public)

	//分类列表
	r.GET("/type", middleware.RoleMiddleWare(), controller.Types)
	//编辑页面
	r.GET("/type/page/:id", middleware.RoleMiddleWare(), controller.EditPage)
	//新增分类
	r.GET("/type/add/:name", middleware.RoleMiddleWare(), controller.TypeAdd)
	//修改分类
	r.GET("/type/edit/:id/:name", middleware.RoleMiddleWare(), controller.TypeEdit)
	//删除分类
	r.GET("/type/delete/:id", middleware.RoleMiddleWare(), controller.DeleteType)
}
