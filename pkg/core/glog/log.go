package glog

import (
	"bufio"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/whilesun/go-admin/pkg/core/gconfig"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type LogConfig struct {
	Type         string
	Path         string
	FileName     string
	MaxAge       int
	RotationTime int
	Stdout       bool
	LogLevel     logrus.Level
}

// getFilePath 定位在pkg同级
func getFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(file) + "../../../..")
	return path
}

func New(logKey string) *logrus.Logger {
	// 默认log配置
	if logKey == "" {
		logKey = "log"
	}
	logConfig := &LogConfig{
		Type:         "file",
		Path:         "runtime/logs",
		FileName:     "sys",
		MaxAge:       7 * 24,
		RotationTime: 24,
		Stdout:       false,
		LogLevel:     logrus.InfoLevel,
	}
	gconfig.Config.UnmarshalKey(logKey, logConfig)
	logConfig.Path = path.Join(getFilePath(), logConfig.Path)
	return logConfig.run()
}

func (logConfig *LogConfig) run() *logrus.Logger {
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
		logConfig.configLocalFileLogger(logger)
	} else {
		log.Fatalf("config logger type [%s] is not support, choose types [file]", logType)
	}
	return logger
}

// ConfigLocalFileLogger 写入文件
/**
logPath logs文件目录
logFileName 文件名
maxAge 文件最大保存时间
rotationTime 日志切割时间
*/
func (logConfig *LogConfig) configLocalFileLogger(log *logrus.Logger) {
	logPath := logConfig.Path
	logFileName := logConfig.FileName
	maxAge := logConfig.MaxAge
	rotationTime := logConfig.RotationTime
	////文件目录
	baseLogPath := path.Join(logPath, logFileName)
	writer, err1 := rotatelogs.New(
		baseLogPath+"_access_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),                               // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),             // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 日志切割时间间隔
	)
	if err1 != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err1))
	}
	writerError, err2 := rotatelogs.New(
		baseLogPath+"_error_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),                               // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),             // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 日志切割时间间隔
	)
	if err2 != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err2))
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.FatalLevel: writerError,
		logrus.ErrorLevel: writerError,
		logrus.PanicLevel: writerError,
	}
	Hook := lfshook.NewHook(writeMap, new(LogFormatter))
	log.AddHook(Hook)
}

//func ConfigESLogger(esUrl string, esHOst string, index stringm) {
//	client, err := elastic.NewClient(elastic.SetURL(esUrl))
//	if err != nil {
//		log.Errorf("config es logger error. %+v", errors.WithStack(err))
//	}
//	esHook, err := elogrus.NewElasticHook(client, esHOst, log.DebugLevel, index)
//	if err != nil {
//		log.Errorf("config es logger error. %+v", errors.WithStack(err))
//	}
//	log.AddHook(esHook)
//}
