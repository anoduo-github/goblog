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
	u := r.Group("", middleware.UserMiddleWare())
	//不需要验证即可访问的
	{
		u.GET("/", controller.Index)             //首页
		u.GET("/user/info", controller.UserInfo) //用户登录信息

		r.GET("/login", controller.LoginPage)        //登录页
		r.POST("/dologin", controller.Login)         //登录
		r.GET("/logout", controller.Logout)          //注销
		r.GET("/captcha", controller.Captcha)        //验证码
		r.GET("/register", controller.RegisterPage)  //注册页
		r.POST("/doregister", controller.Register)   //注册
		r.GET("/footer/info", controller.FooterInfo) //页底信息

		u.GET("/error/:name", controller.Error) //错误

		u.GET("/blog/list/type", controller.TypeBlogList)  //分类博客页
		u.GET("/blog/list/search", controller.SearchBlog)  //模糊查询博客页
		u.GET("/blog/details/:id", controller.BlogDetails) //博客详情
		u.POST("/like/add", controller.AddLike)            //点赞

		u.GET("/archives", controller.Archives)   //时间轴页面
		u.GET("/link", controller.LinkPage)       //友链页面
		u.GET("/picture", controller.PicturePage) //照片墙页面
		u.GET("/about", controller.About)         //关于我页面

		u.GET("/comment/:index", controller.Comments) //留言板页面
		u.POST("/comment/add", controller.AddComment) //添加留言
	}

	//下面是需要进行验证的
	v := r.Group("", middleware.RoleMiddleWare())

	v.GET("/admin", controller.Admin) //管理平台首页

	//博客组
	{
		v.GET("/blog/list", controller.BlogList)                 //博客列表页
		v.GET("/blog/edit", controller.BlogEdit)                 //编辑新博客
		v.GET("/blog/edit/:id", controller.BlogEdit)             //编辑要修改博客
		v.POST("/blog/submit", controller.BlogSubmit)            //提交博客
		v.POST("/blog/upload/picture", controller.UploadPicture) //上传图片
		v.GET("/blog/delete/:id", controller.DeleteBlog)         //删除博客
	}

	//分类组
	{
		v.GET("/type/list/:index", controller.Types)       //分类列表
		v.GET("/type/page/:id", controller.EditPage)       //编辑页面
		v.GET("/type/add/:name", controller.TypeAdd)       //新增分类
		v.GET("/type/edit/:id/:name", controller.TypeEdit) //修改分类
		v.GET("/type/delete/:id", controller.DeleteType)   //删除分类
	}

	//友链组
	{
		v.GET("/link/list/:index", controller.Links)     //友链列表
		v.GET("/link/edit/:id", controller.EditLink)     //友链编辑页面
		v.POST("/link/add", controller.AddLink)          //友链编辑
		v.GET("/link/delete/:id", controller.DeleteLink) //删除友链
	}

	//照片组
	{
		v.GET("/picture/list/:index", controller.Pictures)     //照片列表
		v.GET("/picture/edit/:id", controller.EditPicture)     //照片编辑页面
		v.POST("/picture/add", controller.AddPicture)          //照片编辑
		v.GET("/picture/delete/:id", controller.DeletePicture) //删除照片
	}

	//用户组
	{
		v.GET("/user/list", controller.Users)            //用户列表
		v.GET("/user/detail/:id", controller.UserDetail) //用户详情
		v.GET("/user/delete/:id", controller.DeleteUser) //删除用户
	}

	//留言
	v.GET("/comment/delete/:id", controller.DeleteComment) //删除留言
}
