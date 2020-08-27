package client

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/wufenqiang/MPDCDSPro/src/conf"
	"github.com/wufenqiang/MPDCDSPro/src/logger"
	"github.com/wufenqiang/MPDCDSPro/src/thrift/MPDCDS_BackendService"
	"go.uber.org/zap"
	"net"
	"os"
)

//创建客户端连接，获取连接对象
func Connect() (*MPDCDS_BackendService.MPDCDS_BackendServiceClient, thrift.TTransport) {
	thrifthostport := net.JoinHostPort(conf.Sysconfig.ThriftHost, conf.Sysconfig.ThriftPort)
	logger.GetLogger().Info("Default Thrift Host Port:" + thrifthostport)
	return ConnectHostPort(thrifthostport)
}
func ConnectHostPort(hostport string) (*MPDCDS_BackendService.MPDCDS_BackendServiceClient, thrift.TTransport) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort(conf.Sysconfig.ThriftHost, conf.Sysconfig.ThriftPort))
	if err != nil {
		logger.GetLogger().Error("Get thrift transport failed！", zap.String("error", err.Error()))
	}
	trans := thrift.NewTFramedTransport(transport)
	//useTransport := transportFactory.GetTransport(transport)
	client := MPDCDS_BackendService.NewMPDCDS_BackendServiceClientFactory(trans, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+conf.Sysconfig.ThriftHost, conf.Sysconfig.ThriftPort, " ", err)
		logger.GetLogger().Error("Error opening socket", zap.String("ThriftHost", conf.Sysconfig.ThriftHost), zap.String("ThriftPort", conf.Sysconfig.ThriftPort))
	}
	//defer transport.Close()
	return client, transport
}

//释放transport
func Close(transport thrift.TTransport) {
	defer transport.Close()
}
