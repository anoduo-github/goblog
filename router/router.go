package router

import (
	"fmt"
	"goblog/controller"
	"goblog/module/config"
	"goblog/module/log"
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
	r.LoadHTMLGlob("views/*")

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
	r.GET("/loginpage", controller.LoginPage)
	//登录
	r.GET("/login", controller.Login)
	//验证码
	r.GET("/captcha", controller.Captcha)

	//注册页
	r.GET("/registerpage", controller.RegisterPage)
	//注册
	r.POST("/register", controller.Register)
}

/* func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	var authMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.AllUserAuthorizator{})
	r.POST("/login", authMiddleware.LoginHandler)
	//404 handler
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		code := e.PAGE_NOT_FOUND
		c.JSON(404, gin.H{"code": code, "message": e.GetMsg(code)})
	})
	auth := r.Group("/auth")
	{
		// Refresh time can be longer than token timeout
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	api := r.Group("/user")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/info", v1.GetUserInfo)
		api.POST("/logout", v1.Logout)
	}

	var adminMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.AdminAuthorizator{})
	apiv1 := r.Group("/api/v1")
	//使用AdminAuthorizator中间件，只有admin权限的用户才能获取到接口
	apiv1.Use(adminMiddleware.MiddlewareFunc())
	{
		//vue获取table信息
		apiv1.GET("/table/list", v2.GetArticles)
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	var testMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.TestAuthorizator{})
	apiv2 := r.Group("/api/v2")
	apiv2.Use(testMiddleware.MiddlewareFunc())
	{
		//获取文章列表
		apiv2.GET("/articles", v2.GetArticles)
	}

	return r
} */
