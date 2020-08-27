package logger

import (
	"github.com/wufenqiang/MPDCDSPro/src/conf"
	"github.com/wufenqiang/MPDCDSPro/src/utils"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

//初始化
func init() {
	logger := utils.LoggerConfiger(conf.ProjectName, conf.Sysconfig.LoggerLevel)

	//logger.Info("Logger init......")

	zapLogger = logger
}

//获取zap logger 对象
func GetLogger() *zap.Logger {
	return zapLogger
}
