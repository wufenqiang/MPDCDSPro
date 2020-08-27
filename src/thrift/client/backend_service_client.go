package client

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/wufenqiang/MPDCDSPro/src/thrift/MPDCDS_BackendService"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//创建客户端连接，获取连接对象
//func Connect() (*MPDCDS_BackendService.MPDCDS_BackendServiceClient, thrift.TTransport) {
//	thrifthostport := net.JoinHostPort(conf.Sysconfig.ThriftHost, conf.Sysconfig.ThriftPort)
//	logger.GetLogger().Info("Default Thrift Host Port:" + thrifthostport)
//	return ConnectHostPort(thrifthostport)
//}
func ConnectHostPortLog(zapLogger *zap.Logger, hostport string) (*MPDCDS_BackendService.MPDCDS_BackendServiceClient, thrift.TTransport) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(hostport)
	if err != nil {
		zapLogger.Error("Get thrift transport failed！", zap.String("error", err.Error()))
	}
	trans := thrift.NewTFramedTransport(transport)
	//useTransport := transportFactory.GetTransport(transport)
	client := MPDCDS_BackendService.NewMPDCDS_BackendServiceClientFactory(trans, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+hostport, " ", err)
		zapLogger.Error("Error opening socket", zap.String("ThriftHostPort", hostport))
	}
	//defer transport.Close()
	return client, transport
}
func ConnectHostPort(hostport string) (*MPDCDS_BackendService.MPDCDS_BackendServiceClient, thrift.TTransport) {
	core := zapcore.NewTee()
	zapLogger := zap.New(core)
	return ConnectHostPortLog(zapLogger, hostport)
}

//释放transport
func Close(transport thrift.TTransport) {
	defer transport.Close()
}
