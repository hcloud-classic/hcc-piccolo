package iputil

import (
	"net"
)

// CheckValidIP : Parses string value of IPv4 address then return as net.IP.
// If given wrong IP address, it wil return nil.
func CheckValidIP(ip string) net.IP {
	netIP := net.ParseIP(ip).To4()
	return netIP
}
