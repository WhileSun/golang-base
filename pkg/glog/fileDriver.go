package glog

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

/**
	logPath logs文件目录
	logFileName 文件名
	maxAge 文件最大保存时间
	rotationTime 日志切割时间
 */
func ConfigLocalFileLogger(log *logrus.Logger){
	logPath := logConfig.Path
	logFileName := logConfig.FileName
	maxAge := logConfig.MaxAge
	rotationTime := logConfig.RotationTime
	////文件目录
	baseLogPath := path.Join(logPath,logFileName)
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

