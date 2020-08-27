package main

import (
	"220.243.129.233/wufenqiang/MPDCDSPro/src/logger"
	"220.243.129.233/wufenqiang/MPDCDSPro/src/thrift/client"
	"fmt"
)

func main() {
	logger.GetLogger().Info("test")
	fmt.Println("test")
	client.Connect()
}
