package main

import "github.com/wufenqiang/MPDCDSPro/src/thrift/client"

func main() {
	//fmt.Println("test")
	client.ConnectHostPort("127.0.0.1:19090")

}
