package log

import (
	"github.com/jojopoper/xlog"
)

// LoggerInstance log唯一实例
var LoggerInstance *xlog.XLogger

// NewLoggerInstance 生成log唯一实例
func NewLoggerInstance(file string) *xlog.XLogger {
	ret := new(xlog.XLogger)
	ret.Init(file)
	return ret
}
