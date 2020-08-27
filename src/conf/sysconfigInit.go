package conf

import (
	"github.com/json-iterator/go"
	"github.com/wufenqiang/MPDCDSPro/src/utils"
	"io/ioutil"
	"time"
)

func ReadConf(conffile string, thesysconf interface{}) {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile(conffile)
	if err != nil {
		panic(conffile + "Sys config read err")
	}
	err = jsoniter.Unmarshal(b, thesysconf)
	if err != nil {
		panic(err)
	}
}
func LocalProjectPath() string {
	return utils.ProjectLocation(ProjectName)
}

const Layout = "2006-01-02 15:04:05" //时间格式
const TheLocation = "Asia/Shanghai"

var Loc, _ = time.LoadLocation(TheLocation)
var TimeStamp = time.Now()
