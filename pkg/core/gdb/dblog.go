package gdb

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/whilesun/go-admin/pkg/core/glog"
)

// Writer 定义自己的Writer
type Writer struct {
	log *logrus.Logger
}

// Printf 实现gorm/logger.Writer接口
func (w *Writer) Printf(format string, v ...interface{}) {
	logStr := fmt.Sprintf(format, v...)
	w.log.Info(logStr + "\r\n")
}

// NewWriter 使用glog接管，配置名默认
func NewWriter() *Writer {
	log := glog.New("database.dbLog")
	return &Writer{log: log}
}
