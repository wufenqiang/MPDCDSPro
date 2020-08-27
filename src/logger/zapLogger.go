package logger

import (
	"github.com/wufenqiang/MPDCDSPro/src/conf"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

//初始化
func init() {
	logger := LoggerConfiger(conf.ProjectName, conf.Sysconfig.LoggerLevel)
	zapLogger = logger
}

//获取zap logger 对象
func GetLogger() *zap.Logger {
	return zapLogger
}
