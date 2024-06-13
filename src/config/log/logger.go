package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go.ruoyi.com/src/config/dto"
	"go.ruoyi.com/src/internal/pojo"
	"log"
	"os"
	"time"
)

var (
	Logger   = logrus.New() // 初始化日志对象
	LogEntry *logrus.Entry
)

func Loginit(logConf *pojo.LogCof) {
	//Cfg, _ := ini.Load("config.ini")
	//var (
	//	logPath  = Cfg.Section("client").Key("LogPath").String()
	//	linkName = Cfg.Section("client").Key("LinkName").String()
	//)
	// 写入日志文件
	//logPath := "logs/req/log"         // 日志存放路径
	//linkName := "logs/req/latest.log" // 最新日志的软连接路径

	if _, err := os.Stat(logConf.LogPath); os.IsNotExist(err) {
		if err = os.MkdirAll(logConf.LogPath, 755); err != nil {
			fmt.Println("文件创建错误：", err)
		}
	}
	src, err := os.OpenFile(logConf.LogPath+"/log", os.O_RDWR|os.O_CREATE, 0755) // 初始化日志文件对象
	if err != nil {
		fmt.Println("err: ", err)
	}
	//log := logrus.New()  // 初始化日志对象
	Logger.Out = src // 把产生的日志内容写进日志文件中

	// 日志分隔：1. 每天产生的日志写在不同的文件；2. 只保留一定时间的日志（例如：一星期）
	Logger.SetLevel(logrus.DebugLevel) // 设置日志级别
	logWriter, _ := rotatelogs.New(
		logConf.LogPath+"%Y%m%d.log",              // 日志文件名格式
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 最多保留7天之内的日志
		rotatelogs.WithRotationTime(24*time.Hour), // 一天保存一个日志文件
		rotatelogs.WithLinkName(logConf.LinkName), // 为最新日志建立软连接
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter, // info级别使用logWriter写日志
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 格式日志时间
	})
	Logger.AddHook(Hook)
	log.Printf("日志初始化成功")
	//LogEntry = logrus.NewEntry(Logger).WithField("service", "yi-shou-backstage")
}

func LoggerWithFields(body *dto.LogBody, c *gin.Context) {
	logMap := logrus.Fields{
		"SpendTime": body.SpendTime, //接口花费时间
		"path":      body.Path,      //请求路径
		"Method":    body.Method,    //请求方法
		"Status":    body.Status,    //接口返回状态
		"Proto":     body.Proto,     //http请求版本
		"Ip":        body.Ip,        //IP地址
		"Body":      body.Body,      //请求体
		"Query":     body.Query,     //请求query
		"Message":   body.Message,   //返回错误信息
	}
	Log := Logger.WithFields(logMap)
	if len(c.Errors) > 0 { // 矿建内部错误
		Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
	}
	if body.Status > 200 {
		Log.Error()
	} else {
		Log.Info()
	}
}
