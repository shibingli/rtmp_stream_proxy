package server

import (
	"fmt"
	"io"
	"net"
	pub "rtmp_stream_proxy/commons"
	hdFormat "rtmp_stream_proxy/format"
	"time"

	"github.com/nareix/joy5/av"
	"github.com/nareix/joy5/format"
	"github.com/nareix/joy5/format/rtmp"
	"go.uber.org/zap"
)

type StreamProxy struct {
	RtmpServer *rtmp.Server
}

func NewStreamProxy(srv *rtmp.Server) (*StreamProxy, error) {
	if nil == srv {
		return nil, fmt.Errorf("%s\n", "Invalid parameter")
	}

	return &StreamProxy{RtmpServer: srv}, nil
}

func (s *StreamProxy) HandleConn(c *rtmp.Conn, nc net.Conn) {
	defer func(nc net.Conn) {
		err := nc.Close()
		if err != nil {
			pub.Logger.ZapLogger.Error("HandleConn(Net conn close)", zap.Error(err))
		}
	}(nc)

	if !c.Publishing {
		pub.Logger.ZapLogger.Error("HandleConn(Publishing)",
			zap.Error(fmt.Errorf("%s", "NotPub")),
		)

		return
	}

	fo := hdFormat.NewFormatOpener()

	var err error
	var w *format.Writer

	//配置转发目标
	targetURL := c.URL.String()
	if w, err = fo.Create(targetURL); err != nil {
		if err != io.EOF {
			pub.Logger.ZapLogger.Error("HandleConn(Target Create)", zap.Error(err), zap.String("TargetURL", targetURL))
		}
		return
	}

	c2 := w.Rtmp
	nc2 := w.NetConn

	defer func(nc2 net.Conn) {
		err := nc2.Close()
		if err != nil {
			pub.Logger.ZapLogger.Error("HandleConn(NC2 Close)", zap.Error(err))
		}
	}(nc2)

	pub.Logger.ZapLogger.Info("HandleConn",
		zap.String("Dial", "OK"),
		zap.String("TargetURL", targetURL),
	)

	for {
		var pkt av.Packet

		if pkt, err = c.ReadPacket(); err != nil {
			break
		}

		if err = c2.WritePacket(pkt); err != nil {
			break
		}
	}

	pub.Logger.ZapLogger.Debug("HandleConn(Close)",
		zap.String("RTMP Conn close", "Closed"),
		zap.String("RTMP Target conn close ", "Closed"),
	)
}

func (s *StreamProxy) LogEvent(c *rtmp.Conn, nc net.Conn, e int) {
	es := rtmp.EventString[e]

	pub.Logger.ZapLogger.Debug("LogEvent",
		zap.String("LocalAddr", nc.LocalAddr().String()),
		zap.String("RemoteAddr", nc.RemoteAddr().String()),
		zap.Any("EventString", es),
	)
}

func (s *StreamProxy) Run() error {
	listener, err := net.Listen("tcp", pub.SrvConf.Server.Addr())
	if nil != err {
		return err
	}

	s.RtmpServer.LogEvent = s.LogEvent
	s.RtmpServer.HandleConn = s.HandleConn

	for {
		conn, err := listener.Accept()
		if nil != err {
			pub.Logger.ZapLogger.Warn("Listener accept", zap.Error(err))

			time.Sleep(time.Second)
			continue
		}

		go s.RtmpServer.HandleNetConn(conn)
	}
}
