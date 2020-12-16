package models

import (
	"encoding/binary"
	"net"
	"strings"
)

type AppleIP struct{}

// GetAppleIPs 获取 apple IP
func (ai *AppleIP) GetAppleIPs() (ips []string) {
	ips = cache.SMembers("appleIps")
	return
}

// GetAppleIPs2uint32 获取 apple IP 转换成十进制后的数组
func (ai *AppleIP) GetAppleIPs2uint32() (IPs [][]uint32) {
	ips := ai.GetAppleIPs()
	IPs = make([][]uint32, len(ips))
	for k, ip := range ips {
		appleIPs := strings.Split(ip, "-")
		for _, appleIP := range appleIPs {
			IPs[k] = append(IPs[k], IP2int(appleIP))
		}
	}
	return
}

// IP2int ip string to int
func IP2int(ip string) uint32 {
	if len(net.ParseIP(ip)) == 16 {
		return binary.BigEndian.Uint32(net.ParseIP(ip)[12:16])
	}
	return binary.BigEndian.Uint32(net.ParseIP(ip))
}

// Int2ip int 2 ip
func Int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
