package conf

import (
	"time"
)



type sysconfig struct {
	//日志存储地址
	LoggerPath  string `json:"LoggerPath"`
	LoggerLevel string `json:"LoggerLevel"`
}

const ProjectName= "MPDCDS_BackendService"

const Layout = "2006-01-02 15:04:05" //时间格式
const TheLocation  = "Asia/Shanghai"
var Loc, _ = time.LoadLocation("Asia/Shanghai")
var TimeStamp=time.Now()