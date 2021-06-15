package utils

import (
	"net"
	"os"
	"strings"
)

//IsPhysicalInterfaceDevice 查看网卡是否是物理设备
func IsPhysicalInterfaceDevice(name string) (bool, error) {
	name = strings.TrimSpace(name)

	const (
		netDevicePath        = "/sys/class/net/"
		virtualNetDevicePath = "/devices/virtual/net/"
	)

	filePath := netDevicePath + name
	virtualFilePath := virtualNetDevicePath + name

	fileInfo, err := os.Lstat(filePath)
	if nil != err {
		return false, err
	}

	if (fileInfo.Mode() & os.ModeSymlink) > 0 {
		linkPath, err := os.Readlink(filePath)
		if err != nil {
			return false, err
		}
		if strings.LastIndex(linkPath, virtualFilePath) == -1 {
			return true, nil
		}
	}

	return false, nil
}

//GetIPs 获取当前机器IP地址,filterPhysical为true时，只获取物理络设备IP地址
func GetIPs(filterPhysical ...bool) (ipv4 []string, err error) {

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	physicalInterface := false
	if len(filterPhysical) > 0 {
		physicalInterface = filterPhysical[0]
	}

	var getIPs = func(inf net.Interface) {
		addrs, _ := inf.Addrs()
		for _, address := range addrs {
			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					ipv4 = append(ipv4, ipNet.IP.String())
				}
			}
		}
	}

	for _, inf := range interfaces {
		if (inf.Flags & net.FlagUp) != 0 {
			if physicalInterface {
				ethName := inf.Name
				pEth, err := IsPhysicalInterfaceDevice(ethName)
				if nil != err {
					return nil, err
				}
				if pEth {
					getIPs(inf)
				}
			} else {
				getIPs(inf)
			}
		}
	}
	return
}
