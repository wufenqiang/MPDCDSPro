package conf

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
	ReadConf(conffile, Sysconfig)
}

const ProjectNamePro = "MPDCDSPro"
