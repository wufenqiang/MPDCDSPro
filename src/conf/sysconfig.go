package conf

import (
	"time"
)

type sysconfig struct {

	//thrift 服务ip
	ThriftHost string `json:"ThriftHost"`
	ThriftPort string `json:"ThriftPort"`

	//日志存储地址
	LoggerLevel string `json:"LoggerLevel"`
}

var Sysconfig = &sysconfig{}

func init() {
	dir := LocalProjectPath()

	conffile := dir + "/config.json"

	ReadConf(conffile)
}

const ProjectName = "MPDCDSPro"
const Layout = "2006-01-02 15:04:05" //时间格式
const TheLocation = "Asia/Shanghai"

var Loc, _ = time.LoadLocation(TheLocation)
var TimeStamp = time.Now()
