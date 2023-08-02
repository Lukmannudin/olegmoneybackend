package protocp

import (
	"net"
)

func NetIpToProto(ip net.IP) (string, error) {
	return ip.String(), nil
}

func NetIpFromProto(s string) (net.IP, error) {
	return net.ParseIP(s), nil
}
