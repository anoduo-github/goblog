package log

import (
	"fmt"
	"goblog/module/config"
	"os"

	"github.com/sirupsen/logrus"
)

//Logger 全局log
var Logger *logrus.Logger

//Init 初始化日志模块
func Init() {
	Logger = logrus.New()
	runmode := config.Cfg.Section("log").Key("runmode").String()
	//设置日志等级
	switch runmode {
	case "trace":
		Logger.SetLevel(logrus.TraceLevel)
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		Logger.SetLevel(logrus.FatalLevel)
	case "panic":
		Logger.SetLevel(logrus.PanicLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
	//设置日志格式
	Logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:     false,                 //格式化
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
	})
	//设置日志定位
	Logger.SetReportCaller(true)
	//设置输出位置
	os.MkdirAll("log", os.ModePerm)
	logfile, err := os.OpenFile("log/blog.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file log/blog.log error: ", err)
		os.Exit(1)
	}
	Logger.SetOutput(logfile) //默认为os.stderr
}
