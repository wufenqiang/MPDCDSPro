package conf

import (
	"time"
)

type sysconfig struct {

	//thrift 服务ip
	ThriftHost string `json:"ThriftHost"`
	ThriftPort string `json:"ThriftPort"`

	//日志存储地址
	LoggerPath  string `json:"LoggerPath"`
	LoggerLevel string `json:"LoggerLevel"`

	//日志中显示相关密文
	ShadeInLog bool `json:ShadeInLog`
}

const ProjectName = "MPDCDSPro"

const Layout = "2006-01-02 15:04:05" //时间格式
const TheLocation = "Asia/Shanghai"

var Loc, _ = time.LoadLocation("Asia/Shanghai")
var TimeStamp = time.Now()
