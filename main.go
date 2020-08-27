package main

import (
	"fmt"
	"github.com/wufenqiang/MPDCDSPro/src/logger"
	"github.com/wufenqiang/MPDCDSPro/src/thrift/client"
)

func main() {
	logger.GetLogger().Info("test")
	fmt.Println("test")
	client.Connect()
}
