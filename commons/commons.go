package commons

import (
	gLog "log"
	"rtmp_stream_proxy/config"
	"rtmp_stream_proxy/log"
)

var (
	Logger  *log.Logger
	SrvConf *config.AppConf
)

func Init() {
	srvConf, err := config.InitConfig()
	if nil != err {
		gLog.Fatalf("%s\n", err.Error())
		return
	}

	SrvConf = srvConf

	lev := log.GetLogLevel(SrvConf.App.LogLevel)

	logger, err := log.NewLogger(lev)
	if nil != err {
		gLog.Fatalf("%s\n", err.Error())
		return
	}

	Logger = logger
}
