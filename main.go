package main

import (
	pub "rtmp_stream_proxy/commons"
	"rtmp_stream_proxy/server"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/nareix/joy5/format/rtmp"
)

var (
	once sync.Once
)

func init() {
	once.Do(Init)
}

func Init() {
	pub.Init()
}

func main() {
	sp, err := server.NewStreamProxy(rtmp.NewServer())
	if nil != err {
		pub.Logger.ZapLogger.Fatal("HSP NewStreamProxy", zap.Error(err))
		return
	}

	sp.RtmpServer.HandshakeTimeout = time.Duration(pub.SrvConf.Server.HandshakeTimeout) * time.Second

	if err := sp.Run(); nil != err {
		pub.Logger.ZapLogger.Fatal("HSP Run", zap.Error(err))
	}
}
