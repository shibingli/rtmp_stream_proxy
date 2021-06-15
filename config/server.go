package config

import (
	"fmt"
	"net"
	"strings"
)

type Server struct {
	IP               string `json:"ip" yaml:"ip" mapstructure:"ip"`
	Port             int    `json:"port" yaml:"port" mapstructure:"port"`
	HandshakeTimeout int    `json:"handshake_timeout" yaml:"handshake_timeout" mapstructure:"handshake_timeout"`
}

func (s *Server) Addr() string {
	ip := strings.TrimSpace(s.IP)
	if "" == ip {
		ip = net.IPv4zero.String()
	}
	return fmt.Sprintf("%s:%d", ip, s.Port)
}
