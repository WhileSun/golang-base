package glog

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/whilesun/go-admin/pkg/gconf"
	"log"
	"os"
	"strings"
)

type logConfigObj struct {
	Type         string
	Path         string
	FileName     string
	MaxAge       int
	RotationTime int
	Stdout       bool
	LogLevel     logrus.Level
}

var logConfig *logConfigObj
var glog *logrus.Logger

func Run() {
	if glog != nil{
		return
	}
	logConfig = &logConfigObj{
		Type:         "file",
		Path:         "runtime/logs",
		FileName:     "sys",
		MaxAge:       7 * 24,
		RotationTime: 24,
		Stdout:       false,
		LogLevel:   logrus.InfoLevel,
	}
	gconf.Config.UnmarshalKey("log", logConfig)
	glog = initLog()
}

func Get() *logrus.Logger {
	return glog
}

func initLog() *logrus.Logger {
	//logrus初始化
	logger := logrus.New()
	logger.SetLevel(logConfig.LogLevel)
	logger.SetReportCaller(true)
	//logger标准化日志
	if logConfig.Stdout == true {
		logger.SetFormatter(new(LogFormatter))
		logger.SetOutput(os.Stdout)
		return logger
	}
	//判断日志类型
	logType := strings.ToLower(logConfig.Type)
	if logType == "file" {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatalf("Open Src File err %+v", err)
		}
		writer := bufio.NewWriter(src)
		logger.SetOutput(writer)
		ConfigLocalFileLogger(logger)
	} else {
		log.Fatalf("config logger type [%s] is not support, choose types [file]", logType)
	}
	return logger
}
