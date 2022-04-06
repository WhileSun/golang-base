package glog

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

//DbWriter 定义自己的Writer
type DbWriter struct {
	mlog *logrus.Logger
}

//Printf 实现gorm/logger.Writer接口
func (m *DbWriter) Printf(format string, v ...interface{}) {
	logStr := fmt.Sprintf(format, v...)
	//利用logrus记录日志
	m.mlog.Info(logStr)
}

func NewDbWriter() *DbWriter {
	log := Get()
	return &DbWriter{mlog: log}
}
