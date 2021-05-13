package router

import (
	"fmt"
	"goblog/module/log"
	"goblog/utils/config"
	"os"

	"github.com/gin-gonic/gin"

)

func Run() {
	//1.设置debug模式
	gin.SetMode(gin.DebugMode)

	//2.创建路由
	r := gin.Default()

	//3.设置静态资源文件目录，并且绑定一个Url前缀
	r.Static("static", "static")

	//4.加载templates目录下面的所有模版文件，模版文件扩展名随意
	r.LoadHTMLGlob("views/*")

	//5.初始化路由绑定
	initRouter(r)

	//6.监听端口，启动服务
	port := config.Cfg.Section("blog").Key("port").MustInt(8080)
	err := r.Run(fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Logger.Errorf("gin could not Run error: %v", err)
		os.Exit(1)
	}
}

func initRouter(r *gin.Engine) {
	//首页
	//r.GET("/", controller.Index)
}
