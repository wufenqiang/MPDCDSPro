package conf

import (
	"github.com/json-iterator/go"
	"github.com/wufenqiang/MPDCDSPro/src/utils"
	"io/ioutil"
)

func ReadConf(conffile string) {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile(conffile)
	if err != nil {
		panic(conffile + "Sys config read err")
	}
	err = jsoniter.Unmarshal(b, Sysconfig)
	if err != nil {
		panic(err)
	}
}
func LocalProjectPath() string {
	return utils.ProjectLocation(ProjectName)
}
